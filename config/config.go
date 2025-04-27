package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type ConfigStruct struct {
	Password          string
	BotToken          string
	BotName           string
	DSN               string
	Port              string
	S3ApiUrl          string
	AccessToken       string
	SecretKey         string
	BucketName        string
	WebappName        string
	ClientUrl         string
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
	BotToken := os.Getenv("BOT_TOKEN")
	BotName := os.Getenv("BOT_NAME")
	DSN := os.Getenv("DSN")
	port := os.Getenv("PORT")
	S3ApiUrl := os.Getenv("S3_API_URL")
	BucketName := os.Getenv("BUCKET_NAME")
	WebappName := os.Getenv("WEBAPP_NAME")
	ClientUrl := os.Getenv("CLIENT_URL")
	AccessToken := os.Getenv("ACCESS_TOKEN")
	SecretKey := os.Getenv("SECRET_ACCESS_KEY")
	MongoUrl := os.Getenv("MONGO_URL")
	YookassaAPIURL := os.Getenv("YOOKASSAAPIURL")
	YooKassaSecretKey := os.Getenv("YOOKASSASECRETKEY")
	YookassaShopID := os.Getenv("YOOKASSASHOPID")

	Config = ConfigStruct{
		Password:          Password,
		BotToken:          BotToken,
		BotName:           BotName,
		DSN:               DSN,
		Port:              port,
		S3ApiUrl:          S3ApiUrl,
		AccessToken:       AccessToken,
		SecretKey:         SecretKey,
		BucketName:        BucketName,
		WebappName:        WebappName,
		ClientUrl:         ClientUrl,
		MongoUrl:          MongoUrl,
		YookassaAPIURL:    YookassaAPIURL,
		YooKassaShopID:    YookassaShopID,
		YooKassaSecretKey: YooKassaSecretKey,
	}
}
