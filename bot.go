package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

// BotReturnError returns a message with the error provided
func BotReturnError(bot *tgbotapi.BotAPI, update tgbotapi.Update, err string) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, err)
	bot.Send(msg)
}
