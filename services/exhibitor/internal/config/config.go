package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DatabaseURL      string
	ClerkSecretKey   string
	CLOUDINARY_URL   string
	CLERK_SECRET_KEY string
}

func LoadEnv() *Config {

	if err := godotenv.Load("/home/shivajee/Desktop/opexn-expo/.env"); err != nil {
		log.Println("Warning: No .env file found, using system environment variables")
	}

	return &Config{
		DatabaseURL:      os.Getenv("NEON_DB"),
		ClerkSecretKey:   os.Getenv("CLERK_SECRET_KEY"),
		CLOUDINARY_URL:   os.Getenv("CLOUDINARY_URL"),
		CLERK_SECRET_KEY: os.Getenv("CLERK_SECRET_KEY"),
	}
}
