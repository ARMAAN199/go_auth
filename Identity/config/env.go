package config

import (
	"fmt"
	"os"
	"strconv"
	"time"

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

type RedisConfig struct {
	Host     string
	Port     string
	Password string
	DB       int
}

type AppConfig struct {
	AccessTokenExpiryInMinutes  time.Duration
	RefreshTokenExpiryInMinutes time.Duration
}

var (
	DbConfigs    *Config
	RedisConfigs *RedisConfig
	AppConfigs   *AppConfig
)

func InitConfig() *Config {

	if DbConfigs != nil {
		return DbConfigs
	}

	godotenv.Load()

	DbConfigs = &Config{
		DBHost:     getEnv("DB_HOST", "null"),
		DBUser:     getEnv("DB_USER", "null"),
		DBPassword: getEnv("DB_PASSWORD", "null"),
		DBname:     getEnv("DB_NAME", "null"),
		DBPort:     getEnv("DB_PORT", "null"),
		SSLMode:    getEnv("DB_SSLMODE", "null"),
	}

	return DbConfigs
}

func InitRedisConfig() *RedisConfig {

	if RedisConfigs != nil {
		return RedisConfigs
	}

	godotenv.Load()

	RedisConfigs = &RedisConfig{
		Host:     getEnv("REDIS_HOST", "null"),
		Port:     getEnv("REDIS_PORT", "null"),
		Password: getEnv("REDIS_PASSWORD", "null"),
		DB:       0,
	}

	return RedisConfigs
}

func InitAppConfigs() *AppConfig {

	if AppConfigs != nil {
		return AppConfigs
	}

	godotenv.Load()

	fmt.Print("Loading app configs")

	access, err := strconv.Atoi(getEnv("ACCESS_TOKEN_EXPIRY_IN_MINUTES", "null"))
	if err != nil {
		panic(err)
	}
	refresh, err := strconv.Atoi(getEnv("REFRESH_TOKEN_EXPIRY_IN_MINUTES", "null"))
	if err != nil {
		panic(err)
	}

	AppConfigs = &AppConfig{
		AccessTokenExpiryInMinutes:  time.Duration(access) * time.Minute,
		RefreshTokenExpiryInMinutes: time.Duration(refresh) * time.Minute,
	}

	return AppConfigs
}

func getEnv(key string, fallback string) string {
	value, ok := os.LookupEnv(key)
	if !ok {
		return fallback
	}
	return value
}
