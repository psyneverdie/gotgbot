package main

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

func main() {
	token := "YOUR_TELEGRAM_BOT_TOKEN"
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Fatal(err)
	}

	bot.Debug = true

	log.Printf("Bot %s has been authorized", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)
	if err != nil {
		log.Fatal(err)
	}

	for update := range updates {
		if update.Message == nil { // ignore non-Message updates
			continue
		}

		var responseText string
		var keyboard tgbotapi.ReplyKeyboardMarkup

		// Приветственное сообщение
		if update.Message.IsCommand() && update.Message.Command() == "start" {
			responseText = "Привет! Я твой чат-бот."
			keyboard = tgbotapi.NewReplyKeyboard(
				tgbotapi.NewKeyboardButtonRow(
					tgbotapi.NewKeyboardButton("Кнопка 1"),
					tgbotapi.NewKeyboardButton("Кнопка 2"),
					tgbotapi.NewKeyboardButton("Кнопка 3"),
					tgbotapi.NewKeyboardButton("Кнопка 4"),
				),
			)
		} else {
			switch update.Message.Text {
			case "Кнопка 1", "Кнопка 2", "Кнопка 3", "Кнопка 4":
				responseText = "Вы нажали нажали на кнопку первого слоя — " + update.Message.Text
				keyboard = tgbotapi.NewReplyKeyboard(
					tgbotapi.NewKeyboardButtonRow(
						tgbotapi.NewKeyboardButton("Второй слой: Кнопка 1"),
						tgbotapi.NewKeyboardButton("Второй слой: Кнопка 2"),
					),
					tgbotapi.NewKeyboardButtonRow(
						tgbotapi.NewKeyboardButton("Назад"),
					),
				)
			case "Второй слой: Кнопка 1", "Второй слой: Кнопка 2":
				responseText = "Работа завершена. Вы нажали нажали на кнопку второго слоя — " + update.Message.Text
				// Убираем клавиатуру
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, responseText)
				//msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(false) // опционально
				msg.ReplyMarkup = tgbotapi.ReplyKeyboardRemove{RemoveKeyboard: true}
				_, err := bot.Send(msg)
				if err != nil {
					log.Printf("Error sending message: %v", err)
				}
			case "Назад":
				responseText = "Вы вернулись на первый слой"
				keyboard = tgbotapi.NewReplyKeyboard(
					tgbotapi.NewKeyboardButtonRow(
						tgbotapi.NewKeyboardButton("Кнопка 1"),
						tgbotapi.NewKeyboardButton("Кнопка 2"),
						tgbotapi.NewKeyboardButton("Кнопка 3"),
						tgbotapi.NewKeyboardButton("Кнопка 4"),
					),
				)
			default:
				responseText = "Я не понимаю вашу команду. Пожалуйста, выберите одну из кнопок."
				// Убираем клавиатуру
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, responseText)
				//msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(false) // опционально
				msg.ReplyMarkup = tgbotapi.ReplyKeyboardRemove{RemoveKeyboard: true}
				_, err := bot.Send(msg)
				if err != nil {
					log.Printf("Error sending message: %v", err)
				}
				continue
			}
		}

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, responseText)
		msg.ReplyMarkup = keyboard
		_, err := bot.Send(msg)
		if err != nil {
			log.Printf("Error sending message: %v", err)
		}
	}
}
