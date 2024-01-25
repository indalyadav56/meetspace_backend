package config

import (
	"log"

	"github.com/joho/godotenv"
)

// LoadEnv loads environment variables from a .env file.
func LoadEnv() {
    err := godotenv.Load(".env")
    if err != nil {
        log.Fatal("Error loading .env file:", err)
    }
}