package main

import (
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/vitaly06/geo-weather-bot/internal/config"
	"github.com/vitaly06/geo-weather-bot/internal/config/database"
)

func main() {
	// Load config
	cfg := config.NewConfig()

	// Connect to database
	_, err := database.ConnectDb(cfg.Dsn)

	if err != nil {
		log.Fatal("[ERROR] Database connection error: %w\n", err)
	}

	bot, err := tgbotapi.NewBotAPI(cfg.TelegramBotToken)

	if err != nil {
		log.Fatal("[ERROR] Create bot error: %w\n", err)
	}

	fmt.Println("[LOG] Bot succesfully started")

	bot.Debug = true

	updateConfig := tgbotapi.NewUpdate(0)

	updateConfig.Timeout = 30

	updates := bot.GetUpdatesChan(updateConfig)

	for update := range updates {
		if update.Message == nil {
			continue
		}
		// Echo
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
		msg.ReplyToMessageID = update.Message.MessageID

		if _, err := bot.Send(msg); err != nil {
			log.Fatal("[ERROR] Send message error: %w\n", err)
		}
	}
}
