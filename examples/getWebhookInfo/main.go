package main

import (
	"github.com/nvhai245/kingtalk-bot-api"
	"log"
)

func main() {
	bot, err := ktbotapi.NewBotAPI("1736098:XdTqgOKQCc5oq64IN9TAKdw4qzAnckH1tkV45PGj")
	if err != nil {
		log.Fatal(err)
	}

	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.FirstName)

	l := ktbotapi.NewWebhook("https://webhook.site/6d8aac66-3ade-4fd1-b5f8-63d0a18194bc")
	bot.SetWebhook(l)
	bot.GetWebhookInfo()
}
