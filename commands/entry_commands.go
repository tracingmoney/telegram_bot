package commands

import (
	"fmt"
	"log"
	"strings"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func command_expense(update tgbotapi.Update, bot *tgbotapi.BotAPI) {
	fields := strings.Fields(update.Message.Text)
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
	msg.ParseMode = "markdown"
	defer func() {
		if _, err := bot.Send(msg); err != nil {
			log.Println(err)
		}
	}()

	if len(fields) < 3 {
		msg.Text = "Not enough args for using expense command."
		return
	}

	var count, price, description string
	count = fields[1]
	price = fields[2]
	description = strings.Join(fields[3:], " ")
	date := time.Now()

	msg.Text = fmt.Sprintf("Record added successfully.\nDescription: *%s*\nPrice: *%s*\nCount: *%s*\nDate: *%s*", description, price, count, date.Format(time.ANSIC))
}

func command_income(update tgbotapi.Update, bot *tgbotapi.BotAPI) {
	fields := strings.Fields(update.Message.Text)
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
	msg.ParseMode = "markdown"
	defer func() {
		if _, err := bot.Send(msg); err != nil {
			log.Println(err)
		}
	}()

	if len(fields) < 2 {
		msg.Text = "Not enough args for using income command."
		return
	}

	var income, description string
	income = fields[1]
	description = strings.Join(fields[2:], " ")
	date := time.Now()

	msg.Text = fmt.Sprintf("Record added successfully.\nDescription: *%s*\nIncome: *%s*\nDate: *%s*", description, income, date.Format(time.ANSIC))
}
