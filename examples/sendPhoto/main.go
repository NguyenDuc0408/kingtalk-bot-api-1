package main

import (
	"github.com/nvhai245/kingtalk-bot-api"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

var testChatId int64 = 3337601907588356

func main() {
	bot, err := ktbotapi.NewBotAPI(os.Getenv("KINGTALK_BOT_TOKEN"))
	if err != nil {
		log.Fatal(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)
	absPath, err := filepath.Abs("../image.jpg")
	if err != nil {
	}
	log.Println(err)
	file, err := ioutil.ReadFile(absPath)
	chatMessage := ktbotapi.NewPhotoUpload(testChatId, ktbotapi.FileBytes{
		Name:  "image.jpg",
		Bytes: file,
	})
	_, err = bot.Send(chatMessage)
	if err != nil {
		log.Fatal(err)
	}
}
