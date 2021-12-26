package commands

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/supabase/postgrest-go"
)

func GetCommands() map[string]func(tgbotapi.Update, *tgbotapi.BotAPI, *postgrest.Client) {
	command_hashmap := make(map[string]func(tgbotapi.Update, *tgbotapi.BotAPI, *postgrest.Client))

	// default implemented commands
	command_hashmap["start"] = command_start
	command_hashmap["hi"] = command_hi

	// business logic codes
	command_hashmap["expense"] = command_expense
	command_hashmap["income"] = command_income

	return command_hashmap
}
