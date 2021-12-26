package commands

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/supabase/postgrest-go"
)

func command_start(update tgbotapi.Update, bot *tgbotapi.BotAPI, database *postgrest.Client) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
	msg.Text = `I didn't wrote it for now.`
	if _, err := bot.Send(msg); err != nil {
		log.Println(err)
	}
}

func command_hi(update tgbotapi.Update, bot *tgbotapi.BotAPI, database *postgrest.Client) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Hello !!")
	msg.ReplyToMessageID = update.Message.MessageID

	if _, err := bot.Send(msg); err != nil {
		log.Println(err)
	}
}
