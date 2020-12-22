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

	log.Printf("Authorized on account %s", bot.Self.FirstName)

	c := ktbotapi.BotCommand{
		Command:     "whatIsMyFirstName",
		Description: "I will tell you what your first name is, in case ur mommy didn't.",
	}
	bot.SetCommands([]ktbotapi.BotCommand{c})

	//bot.GetCommands()
}
