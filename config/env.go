package config

import (
	"github.com/joho/godotenv"
	"github.com/phuslu/log"
	"os"
)

type AppEnv struct {
	DbDsn    string
	HttpPort string
}

func InitializeEnv(envFile ...string) *AppEnv {
	err := godotenv.Load(envFile...)
	if err != nil {
		log.Panic().Err(err).Msg("Error loading .env file")
	}
	return &AppEnv{
		DbDsn:    getOrPanic("DB_DSN"),
		HttpPort: getOrDefault("HTTP_PORT", "8080"),
	}
}

func getOrDefault(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

func getOrPanic(key string) string {
	value := os.Getenv(key)
	if value == "" {
		log.Panic().Str("key", key).Msg("Missing environment variable")
	}
	return value
}
