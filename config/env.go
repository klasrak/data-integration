package config

import (
	"os"
)

type Env struct {
	MONGO_URI      string
	JWT_SECRET     string
	ENCRYPT_SECRET string
	REDIS_HOST     string
	REDIS_PORT     string
	REDIS_PASSWORD string
}

func getEnvOrDefault(key, defaultValue string) string {
	value := os.Getenv(key)

	if value == "" {
		return defaultValue
	}

	return value
}

// New creates a new env with environment variables from a .env or with default values
func New() *Env {
	mongoUri := getEnvOrDefault("MONGO_URI", "mongodb://mongo:27017")
	jwtSecret := getEnvOrDefault("JWT_SECRET", "s3cr3ts4uc3")
	encryptSecret := getEnvOrDefault("ENCRYPT_SECRET", "sup3rs3cr3t")
	redisHost := getEnvOrDefault("REDIS_HOST", "redis")
	redisPort := getEnvOrDefault("REDIS_PORT", "6379")
	redisPassword := getEnvOrDefault("REDIS_PASSWORD", "abcd1234")

	return &Env{
		MONGO_URI:      mongoUri,
		JWT_SECRET:     jwtSecret,
		ENCRYPT_SECRET: encryptSecret,
		REDIS_HOST:     redisHost,
		REDIS_PORT:     redisPort,
		REDIS_PASSWORD: redisPassword,
	}
}
