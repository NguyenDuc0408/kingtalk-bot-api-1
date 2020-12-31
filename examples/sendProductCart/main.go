package main

import (
	"github.com/nvhai245/kingtalk-bot-api"
	"log"
)

var testPeerID int32 = 1736117

func main() {
	bot, err := ktbotapi.NewBotAPI("1736098:XdTqgOKQCc5oq64IN9TAKdw4qzAnckH1tkV45PGj")
	if err != nil {
		log.Fatal(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.FirstName)

	msg := ktbotapi.NewProductCartWithPeer(testPeerID, "3FuNpcfFJVcnPGPtWhSVrT941609229347")

	bot.Send(msg)
}
