package tools

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

// Errors to return when failing
var (
	BasicError = "Ha habido un problema al ejecutar el comando."
)

var (
	bot    *tgbotapi.BotAPI
	update *tgbotapi.Update
)

// BotReturnError returns a message with the error provided
func BotReturnError(err string) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, err)
	bot.Send(msg)
}
