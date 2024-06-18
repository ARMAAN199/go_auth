package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBHost     string
	DBUser     string
	DBPassword string
	DBname     string
	DBPort     string
	SSLMode    string
}

func InitConfig() *Config {

	godotenv.Load()

	return &Config{
		DBHost:     getEnv("DB_HOST", "null"),
		DBUser:     getEnv("DB_USER", "null"),
		DBPassword: getEnv("DB_PASSWORD", "null"),
		DBname:     getEnv("DB_NAME", "null"),
		DBPort:     getEnv("DB_PORT", "null"),
		SSLMode:    getEnv("DB_SSLMODE", "null"),
	}
}

func getEnv(key string, fallback string) string {
	value, ok := os.LookupEnv(key)
	if !ok {
		return fallback
	}
	return value
}
