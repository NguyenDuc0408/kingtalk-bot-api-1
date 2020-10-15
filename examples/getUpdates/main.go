package main

import (
	"github.com/nvhai245/kingtalk-bot-api"
	"log"
	"os"
)

func main() {
	bot, err := ktbotapi.NewBotAPI(os.Getenv("KINGTALK_BOT_TOKEN"))
	if err != nil {
		log.Fatal(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	updates, err := bot.GetUpdates(ktbotapi.UpdateConfig{
		Offset:  0,
		Limit:   5,
		Timeout: 5,
	})
	if err != nil {
		log.Fatal(err)
	}
	_ = updates
}
