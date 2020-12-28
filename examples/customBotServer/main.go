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

	u := ktbotapi.NewUpdate(0)
	u.Timeout = 60
	updates, err := bot.GetUpdatesChan(u)
	log.Println("Listening to bot updates...")

	for update := range updates {
		if update.Message == nil {
			if update.InlineQuery == nil {
				continue
			}
			var results []interface{}
			result := ktbotapi.InlineQueryResultPhoto{
				Type:        "photo",
				ID:          "testIDForCustomBotHandler",
				URL:         "https://images-cdn.9gag.com/photo/aq1WXGP_700b.jpg",
				MimeType:    "image/jpg",
				Title:       "yo",
				Description: "",
				Caption:     "yo",
			}
			results = append(results, result)
			inline := ktbotapi.InlineConfig{
				InlineQueryID: update.InlineQuery.ID,
				Results:       results,
			}
			bot.AnswerInlineQuery(inline)
		} else if !update.Message.IsCommand() {

			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

			msg := ktbotapi.NewMessage(update.Message.Chat.ID, "Tại sao tôi phải trả lời bạn?")
			url := "https://www.instagram.com"
			msg.ReplyMarkup = ktbotapi.InlineKeyboardMarkup{InlineKeyboard: [][]ktbotapi.InlineKeyboardButton{
				{ktbotapi.InlineKeyboardButton{
					Text: "Tôi không biết",
					URL:  &url,
				}},
			}}
			bot.Send(msg)
		} else {
			var msg ktbotapi.MessageConfig
			switch update.Message.Command() {
			case "allinlinebuttons":
				msg = ktbotapi.NewMessage(update.Message.Chat.ID, "Các loại inline buttons: ")
				url := "https://www.google.com"

				bt1 := ktbotapi.NewInlineKeyboardButtonURL("Button url", url)
				bt2 := ktbotapi.NewInlineKeyboardButtonSwitch("Button switch", "@hainv_bot")
				bt3 := ktbotapi.NewInlineKeyboardButtonData("Button Data", "ThisIsData")
				bt4 := ktbotapi.NewInlineKeyboardButtonPay("Button Pay")

				msg.ReplyMarkup = ktbotapi.InlineKeyboardMarkup{InlineKeyboard: [][]ktbotapi.InlineKeyboardButton{
					{
						bt1,
						bt2,
						bt3,
						bt4,
					},
				}}
			case "allreplybuttons":
				msg = ktbotapi.NewMessage(update.Message.Chat.ID, "Các loại reply buttons: ")
				bt5 := ktbotapi.NewKeyboardButton("Button text")
				bt6 := ktbotapi.NewKeyboardButtonContact("Request contact")
				bt7 := ktbotapi.NewKeyboardButtonLocation("Request location")
				msg.ReplyMarkup = ktbotapi.ReplyKeyboardMarkup{
					Keyboard: [][]ktbotapi.KeyboardButton{
						{
							bt5,
							bt6,
							bt7,
						},
					},
					ResizeKeyboard:  false, // optional
					OneTimeKeyboard: false, // optional
					Selective:       false, // optional
				}
			default:
				msg.Text = "I don't know that command"
			}
			bot.Send(msg)
		}
	}
}
