package main

import (
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/vitaly06/geo-weather-bot/internal/config"
	"github.com/vitaly06/geo-weather-bot/internal/database"
	"github.com/vitaly06/geo-weather-bot/internal/domain"
	"github.com/vitaly06/geo-weather-bot/internal/repository"
)

var menuKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Указать город"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Профиль"),
		tgbotapi.NewKeyboardButton("Узнать погоду"),
	),
)

var waitingForCity = make(map[int64]bool)

func main() {
	// Load config
	cfg := config.NewConfig()

	// Connect to database
	db, err := database.ConnectDb(cfg.Dsn)

	if err != nil {
		log.Printf("[ERROR] Database connection error: %s\n", err)
	}

	err = database.AutoMigrate(db)

	if err != nil {
		log.Printf("[ERROR] Automigrate error: %s\n", err)
	}

	userR := repository.NewUserRepository(db)

	bot, err := tgbotapi.NewBotAPI(cfg.TelegramBotToken)

	if err != nil {
		log.Printf("[ERROR] Create bot error: %s\n", err)
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

		if update.Message.IsCommand() {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
			switch update.Message.Command() {
			case "start":
				telegramUserID := update.Message.From.ID
				msg.Text = "Добро пожаловать в бота погоды!"
				userR.Create(&domain.User{
					ID: telegramUserID,
				})
				msg.ReplyMarkup = menuKeyboard
			}
			if _, err := bot.Send(msg); err != nil {
				log.Printf("[ERROR] Send message error: %s\n", err)
			}

			continue
		}

		if update.Message.Text == "Указать город" {
			userId := update.Message.From.ID

			waitingForCity[userId] = true

			msg := tgbotapi.NewMessage(
				update.Message.Chat.ID, "Введите название вашего города:",
			)

			if _, err := bot.Send(msg); err != nil {
				log.Printf("[ERROR] Send message error: %s\n", err)
			}

			continue
		}

		// Указание города
		userID := update.Message.From.ID

		if waitingForCity[userID] {
			city := update.Message.Text

			user, err := userR.GetByID(userID)

			if err != nil {
				log.Printf("[ERROR] Get user error: %s", err)
				continue
			}

			user.CityName = city

			if err := userR.Update(user); err != nil {
				log.Printf("[ERROR] Update user error: %s", err)
				continue
			}
			// Удаляем состояние смены города
			delete(waitingForCity, userID)

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, fmt.Sprintf("Город сохранён: %s", city))

			if _, err := bot.Send(msg); err != nil {
				log.Printf("[ERROR] Send message error: %s\n", err)
			}

			continue
		}

		// Echo
		// msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
		// msg.ReplyToMessageID = update.Message.MessageID

	}
}
