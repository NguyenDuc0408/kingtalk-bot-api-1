package main

import (
	"encoding/json"
	"github.com/nvhai245/kingtalk-bot-api"
	"log"
)

var testChatId int64 = 7455805529554338
var testPeerID int32 = 1736098

func main() {
	bot, err := ktbotapi.NewBotAPI("332000:YHE508a4V8HL8yKjlHwchA6RXYEjLpIIXExUzwaH")
	if err != nil {
		log.Fatal(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.FirstName)

	//absPath, err := filepath.Abs("../image.jpg")
	//if err != nil {
	//	log.Println(err)
	//}
	//file, err := ioutil.ReadFile(absPath)
	chatMessage := ktbotapi.NewPhotoShareWithPeer(testPeerID, "2195955903245319218")

	//chatMessage := ktbotapi.NewPhotoUrl(testChatId, "https://i.kym-cdn.com/entries/icons/mobile/000/034/408/Punching_Pepe_Banner.jpg")

	res, err := bot.Send(chatMessage)
	if err != nil {
		log.Fatal(err)
	}
	msg, _ := json.Marshal(res)
	log.Println(string(msg))
}
