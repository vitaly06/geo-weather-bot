package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	TelegramBotToken string
}

func NewConfig() *Config {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return &Config{
		TelegramBotToken: LoadEnv("TELEGRAM_BOT_TOKEN", "no value"),
	}
}

func LoadEnv(key, replacement string) string {
	value := os.Getenv(key)

	if value == "" {
		return replacement
	}

	return value
}
