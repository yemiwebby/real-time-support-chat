package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("Error loading .env file")
	}

	os.Setenv("TWILIO_ACCOUNT_SID", os.Getenv("TWILIO_ACCOUNT_SID"))
	os.Setenv("TWILIO_AUTH_TOKEN", os.Getenv("TWILIO_AUTH_TOKEN"))
	os.Setenv("TWILIO_PHONE_NUMBER", os.Getenv("TWILIO_PHONE_NUMBER"))
}