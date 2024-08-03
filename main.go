package main

import (
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	bot, err := tgbotapi.NewBotAPI(os.Getenv("SPILLATORE_TOKEN"))
	if err != nil {
		log.Fatal(err)
	}

	bot.Debug = true
}
