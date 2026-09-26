package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	TelegramBotToken string
	Dsn              string
}

func NewConfig() *Config {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return &Config{
		TelegramBotToken: LoadEnv("TELEGRAM_BOT_TOKEN", "no value"),
		Dsn:              LoadEnv("DSN", "host=localhost user=your_username password=your_strong_password dbname=geo_weather_db port=5432"),
	}
}

func LoadEnv(key, replacement string) string {
	value := os.Getenv(key)

	if value == "" {
		return replacement
	}

	return value
}
