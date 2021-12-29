package commands

import (
	"fmt"
	"log"
	"strings"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/supabase/postgrest-go"
)

func command_expense(update tgbotapi.Update, bot *tgbotapi.BotAPI, database *postgrest.Client) {
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

	var piece, price, description string
	piece = fields[1]
	price = fields[2]
	description = strings.Join(fields[3:], " ")
	date := time.Now()

	fmt.Println(database.From("expense").Insert(map[string]interface{}{
		"piece_count": piece,
		"price":       price,
		"description": description,
		"date":        date,
	}, false, "", "", "").ExecuteString())

	msg.Text = fmt.Sprintf("Record added successfully.\nDescription: *%s*\nPrice: *%s*\nCount: *%s*\nDate: *%s*", description, price, piece, date.Format(time.ANSIC))
}

func command_income(update tgbotapi.Update, bot *tgbotapi.BotAPI, database *postgrest.Client) {
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

	database.From("income").Insert(map[string]interface{}{
		"income":      income,
		"description": description,
		"date":        date,
	}, false, "", "", "").ExecuteString()

	msg.Text = fmt.Sprintf("Record added successfully.\nDescription: *%s*\nIncome: *%s*\nDate: *%s*", description, income, date.Format(time.ANSIC))
}
