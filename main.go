package main

import (
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/supabase/postgrest-go"
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

	client := postgrest.NewClient(os.Getenv("SUPABASE_URL"), "", map[string]string{}).TokenAuth(os.Getenv("SUPABASE_PUBLIC_KEY"))
	if client.ClientError != nil {
		log.Panic(err)
	}

	command_hashmap := commands.GetCommands()

	updates := bot.GetUpdatesChan(u)
	for update := range updates {
		if update.Message == nil || !update.Message.IsCommand() { // ignore any non-Message updates
			continue
		}

		if handle_func, ok := command_hashmap[update.Message.Command()]; ok {
			handle_func(update, bot, client)
		}
	}
}
