package utils

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadENV(envStr string) string {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	secret := os.Getenv(envStr)
	if secret == "" {
		log.Fatal("error loading secret from .env")
	}

	return secret
}
