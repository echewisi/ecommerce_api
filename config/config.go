package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DB     string
	JWTKey string
}

func LoadConfig() Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return Config{
		DB:     os.Getenv("DATABASE_URL"),
		JWTKey: os.Getenv("JWT_SECRET"),
	}
}
