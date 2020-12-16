package main

import (
	"github.com/nvhai245/kingtalk-bot-api"
	"log"
)

var testChatId int64 = 3337584726760808

func main() {
	bot, err := ktbotapi.NewBotAPI("777576:CCFB1cLCeOeGRdkTLm2uFUtTpLG9vU8QFYGTi2PX")
	if err != nil {
		log.Fatal(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	var entities []ktbotapi.MessageEntity
	entities = append(entities, ktbotapi.MessageEntity{
		Type:   "mention",
		Offset: 0,
		Length: 10,
		URL:    "",
		User:   nil,
	})

	msg := ktbotapi.MessageConfig{
		BaseChat: ktbotapi.BaseChat{
			ChatID:              testChatId,
			ChannelUsername:     "",
			ReplyToMessageID:    0,
			ReplyMarkup:         nil,
			DisableNotification: false,
		},
		Text:                  "@hainv_bot",
		ParseMode:             "",
		DisableWebPagePreview: false,
		Entities:              entities,
	}
	bot.Send(msg)
}
