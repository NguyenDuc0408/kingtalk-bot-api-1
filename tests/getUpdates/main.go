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

	//_, err = bot.SetWebhook(ktbotapi.NewWebhookWithCert("https://www.google.com:8443/"+bot.Token, "cert.pem"))
	//if err != nil {
	//	log.Fatal(err)
	//}
	updates, err := bot.GetUpdates(ktbotapi.UpdateConfig{
		Offset:  0,
		Limit:   5,
		Timeout: 0,
	})
	if err != nil {
		log.Fatal(err)
	}
	_ = updates
	//updates := bot.ListenForWebhook("/" + bot.Token)
	//go http.ListenAndServeTLS("0.0.0.0:8443", "cert.pem", "key.pem", nil)
	//
	//for update := range updates {
	//	log.Printf("%+v\n", update)
	//}
}
