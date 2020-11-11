package main

import (
	"encoding/json"
	"fmt"
	ktbotapi "github.com/nvhai245/kingtalk-bot-api"
	"log"
	"net/http"
)

// This handler is called everytime telegram sends us a webhook event
func Handler(res http.ResponseWriter, req *http.Request) {
	// First, decode the JSON response body
	body := &ktbotapi.Update{}
	if err := json.NewDecoder(req.Body).Decode(body); err != nil {
		fmt.Println("could not decode request body", err)
		return
	}

	// Nếu user gửi tin nhắn cho bot
	if body.Message != nil {
		chatId := body.Message.Chat.ID
		log.Println("chatId: ", chatId)
		if err := taiSaoToiPhaiTraLoiBan(chatId); err != nil {
			fmt.Println("error in sending reply:", err)
			return
		}
		// log a confirmation message if the message is sent successfully
		fmt.Println("reply sent")
	}
}

//The below code deals with the process of sending a response message
// to the user

// taiSaoToiPhaiTraLoiBan takes a chatID and sends "Tại sao tôi phải trả lời bạn?" to them
func taiSaoToiPhaiTraLoiBan(chatID int64) error {
	chatMessage := ktbotapi.NewMessage(chatID, "Tại sao tôi phải trả lời bạn?")
	message, err := bot.Send(chatMessage)
	if err != nil {
		return err
	}
	log.Printf("%+v", message)
	return nil
}

var bot, err = ktbotapi.NewBotAPI("1735940:93M5fKuPedrQ45PiosUTvl0Z0h5lTrCaw6tYRA9a")

// FInally, the main funtion starts our server on port 8080
func main() {
	if err != nil {
		log.Fatal(err)
	}
	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)
	http.ListenAndServe(":3000", http.HandlerFunc(Handler))
}
