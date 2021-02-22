package main

import (
	"encoding/json"
	ktbotapi "github.com/nvhai245/kingtalk-bot-api"
	"log"
)

var testChatId int64 = 7453585031462049

func main() {
	bot, err := ktbotapi.NewBotAPI("1735423:yvisjjtQQmfK91Mv4JlH0UClB6Dt1jtBZak4Tvba")
	if err != nil {
		panic(err)
	}

	bot.Debug = true

	chatMessage := ktbotapi.NewPhotoUrl(testChatId, "http://static10.zamba.vn/static/rdchat/media/2021/01/05/xsu6DdPA4u_1609814736.jpg")

	res, err := bot.Send(chatMessage)
	if err != nil {
		msg, _ := json.Marshal(res)
		log.Print(string(msg))
		panic(err)
	}

	log.Print("ssss")

	msg, _ := json.Marshal(res)
	log.Print(string(msg))
}
