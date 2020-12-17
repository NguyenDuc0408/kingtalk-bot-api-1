package main

import (
	"encoding/json"
	"fmt"
	"github.com/nvhai245/kingtalk-bot-api"
)

func main() {
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
	menuS, _ := json.Marshal(menu)
	fmt.Print(string(menuS))
}
