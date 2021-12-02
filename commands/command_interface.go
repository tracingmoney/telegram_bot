package commands

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func GetCommands(bot *tgbotapi.BotAPI) map[string]func(tgbotapi.Update, *tgbotapi.BotAPI) {
	command_hashmap := make(map[string]func(tgbotapi.Update, *tgbotapi.BotAPI))

	// default implemented commands
	command_hashmap["start"] = command_start
	command_hashmap["hi"] = command_hi

	// business logic codes
	command_hashmap["expense"] = command_expense
	command_hashmap["income"] = command_income

	return command_hashmap
}
