package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type ConfigStruct struct {
	Password          string
	DSN               string
	Port              string
	S3ApiUrl          string
	BucketName        string
	WebappName        string
	MongoUrl          string
	YookassaAPIURL    string
	YooKassaSecretKey string
	YooKassaShopID    string
}

var Config ConfigStruct

func GetConfig() {
	// if err := godotenv.Load("../../.env"); err != nil {
	if err := godotenv.Load(); err != nil {
		log.Printf(".env file not found: %s", err)
	}

	// Getting values from environment file
	Password := os.Getenv("PASSWORD")
	DSN := os.Getenv("DSN")
	port := os.Getenv("PORT")
	S3ApiUrl := os.Getenv("S3_API_URL")
	BucketName := os.Getenv("BUCKET_NAME")
	WebappName := os.Getenv("WEBAPP_NAME")
	MongoUrl := os.Getenv("MONGO_URL")
	YookassaAPIURL := os.Getenv("YOOKASSAAPIURL")
	YooKassaSecretKey := os.Getenv("YOOKASSASECRETKEY")
	YookassaShopID := os.Getenv("YOOKASSASHOPID")

	Config = ConfigStruct{
		Password:          Password,
		DSN:               DSN,
		Port:              port,
		S3ApiUrl:          S3ApiUrl,
		BucketName:        BucketName,
		WebappName:        WebappName,
		MongoUrl:          MongoUrl,
		YookassaAPIURL:    YookassaAPIURL,
		YooKassaShopID:    YookassaShopID,
		YooKassaSecretKey: YooKassaSecretKey,
	}
}
