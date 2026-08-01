package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	AppName string
	Port    string
	Env     string
}

func Load() *Config {

	if err := godotenv.Load(); err != nil {
		log.Println("⚠️ .env not found, using system environment")
	}

	return &Config{
		AppName: getEnv("APP_NAME", "KongPay"),
		Port:    getEnv("PORT", "8080"),
		Env:     getEnv("APP_ENV", "development"),
	}
}

func getEnv(key, fallback string) string {

	value := os.Getenv(key)

	if value == "" {
		return fallback
	}

	return value
}
