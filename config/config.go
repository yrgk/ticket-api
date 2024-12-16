package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type ConfigStruct struct {
	DSN        string
	Port       string
	S3ApiUrl   string
	BucketName string
}

var Config ConfigStruct

func GetConfig() {
	// if err := godotenv.Load("../../.env"); err != nil {
	if err := godotenv.Load(); err != nil {
		log.Printf(".env file not found: %s", err)
	}

	// Getting values from environment file
	DSN := os.Getenv("DSN")
	port := os.Getenv("PORT")
	S3ApiUrl := os.Getenv("S3_API_URL")
	BucketName := os.Getenv("BUCKET_NAME")

	Config = ConfigStruct{
		DSN:  DSN,
		Port: port,
		S3ApiUrl: S3ApiUrl,
		BucketName: BucketName,
	}
}
