package config

import (
	"os"

	"github.com/joho/godotenv"
)

type env struct {
	MONGO_URI      string
	JWT_SECRET     string
	ENCRYPT_SECRET string
}

func getEnvOrDefault(key, defaultValue string) string {
	value := os.Getenv(key)

	if value == "" {
		return defaultValue
	}

	return value
}

// New creates a new env with environment variables from a .env or with default values
func New() *env {
	err := godotenv.Load()

	if err != nil {
		panic(err.Error())
	}

	mongoUri := getEnvOrDefault("MONGO_URI", "mongodb:mongo:27017")
	jwtSecret := getEnvOrDefault("JWT_SECRET", "s3cr3ts4uc3")
	encryptSecret := getEnvOrDefault("ENCRYPT_SECRET", "sup3rs3cr3t")

	return &env{
		MONGO_URI:      mongoUri,
		JWT_SECRET:     jwtSecret,
		ENCRYPT_SECRET: encryptSecret,
	}
}
