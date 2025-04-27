package utils

import (
	"crypto/md5"
	"encoding/hex"

	"bytes"
	"context"
	"fmt"
	"image/png"
	Config "ticket-api/config"
	"ticket-api/internal/models"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/aws/aws-sdk-go-v2/config"

	"github.com/skip2/go-qrcode"
)

func GetMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

func DeleteFromS3(key string) error {
	return nil
}

// func CreateQrCode(body models.TakeTicketRequest, ticketId string, objectName string) (string, error) {
func CreateQrCode(body models.TakeTicketRequest, ticketId string) (string, error) {
	url := fmt.Sprintf("%s?startapp=check_%s", Config.Config.WebappName, ticketId)
	qr, err := qrcode.New(url, qrcode.Medium)
	if err != nil {
		return "", fmt.Errorf("qr generation failed: %w", err)
	}

	// Encoding QR-code
	var buf bytes.Buffer
	err = png.Encode(&buf, qr.Image(256))
	if err != nil {
		return "", fmt.Errorf("png encoding failed: %w", err)
	}

	// Making config
	resolver := aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
		if service == s3.ServiceID && region == "ru-central1" {
			return aws.Endpoint{
				URL:           "https://storage.yandexcloud.net",
				SigningRegion: "ru-central1",
			}, nil
		}
		return aws.Endpoint{}, fmt.Errorf("unknown endpoint requested")
	})

	// Load AWS config with custom resolver
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion("ru-central1"),
		config.WithCredentialsProvider(
			credentials.NewStaticCredentialsProvider(Config.Config.AccessToken, Config.Config.SecretKey, ""),
		),
		config.WithEndpointResolverWithOptions(resolver),
	)
	if err != nil {
		return "", fmt.Errorf("aws config load failed: %w", err)
	}

	// Initialize S3 client
	client := s3.NewFromConfig(cfg)

	// Upload file
	_, err = client.PutObject(context.Background(), &s3.PutObjectInput{
		Bucket:      aws.String(Config.Config.BucketName),
		Key:         aws.String(ticketId),
		Body:        bytes.NewReader(buf.Bytes()),
		ContentType: aws.String("image/png"),
		ACL:         types.ObjectCannedACLPublicRead,
	})
	if err != nil {
		return "", fmt.Errorf("upload failed: %w", err)
	}

	link := fmt.Sprintf("%s/%s/%s", Config.Config.S3ApiUrl, Config.Config.BucketName, ticketId)

	return link, nil
}
