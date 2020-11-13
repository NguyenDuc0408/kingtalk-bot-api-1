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

	u := ktbotapi.NewUpdate(0)
	u.Timeout = 60
	updates, err := bot.GetUpdatesChan(u)
	log.Println("Listening to bot updates...")

	for update := range updates {
		if update.Message == nil {
			continue
		}

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		msg := ktbotapi.NewMessage(update.Message.Chat.ID, "Tại sao tôi phải trả lời bạn?")
		//msg.ReplyToMessageID = update.Message.MessageID
		//btn := ktbotapi.KeyboardButton{
		//	RequestLocation: true,
		//	Text:            "Gimme where u live!!",
		//}
		url := "https://www.instagram.com"
		msg.ReplyMarkup = ktbotapi.InlineKeyboardMarkup{InlineKeyboard: [][]ktbotapi.InlineKeyboardButton{
			{ktbotapi.InlineKeyboardButton{
				Text: "Tôi không biết",
				URL:  &url,
			}},
		}}
		bot.Send(msg)
	}
}
