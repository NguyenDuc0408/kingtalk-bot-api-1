package main

import (
	"github.com/nvhai245/kingtalk-bot-api"
	"log"
	"math/rand"
)

var testChatId int64 = 7455805529554338

func main() {
	bot, err := ktbotapi.NewBotAPI("1736098:XdTqgOKQCc5oq64IN9TAKdw4qzAnckH1tkV45PGj")
	if err != nil {
		log.Fatal(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.FirstName)

	msg := ktbotapi.NewProduct(testChatId, rand.Int63(), "asdjldsgjdsijflakdgjsdijflsajfldskjf")

	bot.Send(msg)
}
