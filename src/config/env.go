package config

import (
	"github.com/joho/godotenv"
	"github.com/phuslu/log"
	"os"
)

type AppEnv struct {
	DbDsn    string
	HttpPort string
	AppHost  string
	Env      string
	AppName  string
}

func (a AppEnv) IsDevelopment() bool {
	return a.Env == "development"
}

func InitializeEnv(envFile ...string) *AppEnv {
	err := godotenv.Load(envFile...)
	if err != nil {
		log.Panic().Err(err).Msg("Error loading .env file")
	}
	return &AppEnv{
		DbDsn:    getOrPanic("DB_DSN"),
		HttpPort: getOrDefault("API_HTTP_PORT", "8080"),
		AppHost:  getOrPanic("APP_HOST"),
		Env:      getOrDefault("APP_ENV", "development"),
		AppName:  getOrDefault("APP_NAME", "fyndr"),
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
