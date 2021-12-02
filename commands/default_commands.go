package commands

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func command_start(update tgbotapi.Update, bot *tgbotapi.BotAPI) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
	msg.Text = `I didn't wrote it for now.`
	if _, err := bot.Send(msg); err != nil {
		log.Println(err)
	}
}

func command_hi(update tgbotapi.Update, bot *tgbotapi.BotAPI) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Hello !!")
	msg.ReplyToMessageID = update.Message.MessageID

	if _, err := bot.Send(msg); err != nil {
		log.Println(err)
	}
}
