package main

import (
	"github.com/nvhai245/kingtalk-bot-api"
	"log"
	"os"
)

var testChatId int64 = 7455805529554338

func main() {
	bot, err := ktbotapi.NewBotAPI(os.Getenv("KINGTALK_BOT_TOKEN"))
	if err != nil {
		log.Fatal(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	chatMessage := ktbotapi.NewMessage(testChatId, "sent by ktbotapi :)")
	message, err := bot.Send(chatMessage)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%+v", message)
}
