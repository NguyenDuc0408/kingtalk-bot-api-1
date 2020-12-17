package main

import (
	"github.com/nvhai245/kingtalk-bot-api"
	"log"
)

func main() {
	bot, err := ktbotapi.NewBotAPI("1736117:2KuXmm6GHWjzbvzOqv4TVxsPoNhmJ3jDd4hgHmFZ")
	if err != nil {
		log.Fatal(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.FirstName)

	uRL := "https://www.facebook.com"
	inlineRow := ktbotapi.NewInlineKeyboardRow(&ktbotapi.InlineKeyboardButton{
		Text: "Facebook",
		URL:  &uRL,
	}, &ktbotapi.InlineKeyboardButton{
		Text: "Second",
	})
	replyRow := ktbotapi.NewKeyboardButtonRow(&ktbotapi.KeyboardButton{
		Text:            "reply with geo",
		RequestLocation: true,
	})
	menu := ktbotapi.UserMenu{
		ReplyKeyBoardButtonRow:  replyRow,
		InlineKeyboardButtonRow: inlineRow,
	}
	bot.SetUserMenu(ktbotapi.UserMenuConfig{Menu: menu})

	//bot.DeleteUserMenu()
}
