package main

import (
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/yusufpapurcu/telegram_bot/commands"
)

func main() {
	log.SetFlags(log.Lshortfile)

	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_APITOKEN"))
	if err != nil {
		log.Panic(err)
	}

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	command_hashmap := commands.GetCommands(bot)

	updates := bot.GetUpdatesChan(u)
	for update := range updates {
		if update.Message == nil || !update.Message.IsCommand() { // ignore any non-Message updates
			continue
		}

		if handle_func, ok := command_hashmap[update.Message.Command()]; ok {
			handle_func(update, bot)
		}
	}
}
