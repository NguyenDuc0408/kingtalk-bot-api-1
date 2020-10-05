# Bindings bằng ngôn ngữ Go cho Bot API của KingTalk

[![GoDoc](https://godoc.org/github.com/go-telegram-bot-api/telegram-bot-api?status.svg)](http://godoc.org/github.com/go-telegram-bot-api/telegram-bot-api)
[![Travis](https://travis-ci.org/go-telegram-bot-api/telegram-bot-api.svg)](https://travis-ci.org/go-telegram-bot-api/telegram-bot-api)

Các method được dựa trên library telegram-bot-api dành cho telegram. Mọi thông tin khác về thư viện có thể được tìm thấy ở [godoc](http://godoc.org/github.com/go-telegram-bot-api/telegram-bot-api) này.

Mục đích của thư viện này là tạo ra một wrapper xung quanh Bot API của KingTalk mà không bổ sung thêm bất cứ tính năng nào. Hiện nay có 
nhiều thư viện khác cũng cung cấp các tính năng giúp tùy biến plugins và xử lý commands cho Bot API, có thể được
sử dụng chung với thư viện này.

Gia nhập [nhóm phát triển Telegram Bot Api](https://telegram.me/go_telegram_bot_api) để đặt câu hỏi hoặc thảo luật thêm về việc phát triển project.

## Example

Đầu tiên, đảm bảo library đã được cài đặt và cập nhận phiên bản mới nhất:

`go get -u github.com/nvhai245/kingtalk-bot-api`.

Ví dụ dưới đây là một bot đơn giản thực hiện lấy về các updates mới nhất rồi đăng chúng vào trong nhóm chat hiện tại:

```go
package main

import (
	"log"

	"github.com/go-telegram-bot-api/telegram-bot-api"
)

func main() {
	bot, err := tgbotapi.NewBotAPI("Token_Của_Bot")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil { // ignore any non-Message Updates
			continue
		}

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
		msg.ReplyToMessageID = update.Message.MessageID

		bot.Send(msg)
	}
}
```

Xem nhiều hơn các ví dụ tại [wiki](https://github.com/go-telegram-bot-api/telegram-bot-api/wiki)
với nhiều thông tin chi tiết hơn vè cách sử dụng api, commands và các reply markup.

Dưới đây là một ví dụ cho bot muốn sử dụng webhook của Google App Engine.

```go
package main

import (
	"log"
	"net/http"

	"github.com/go-telegram-bot-api/telegram-bot-api"
)

func main() {
	bot, err := tgbotapi.NewBotAPI("Token_Của_Bot")
	if err != nil {
		log.Fatal(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	_, err = bot.SetWebhook(tgbotapi.NewWebhookWithCert("https://www.google.com:8443/"+bot.Token, "cert.pem"))
	if err != nil {
		log.Fatal(err)
	}
	info, err := bot.GetWebhookInfo()
	if err != nil {
		log.Fatal(err)
	}
	if info.LastErrorDate != 0 {
		log.Printf("Telegram callback failed: %s", info.LastErrorMessage)
	}
	updates := bot.ListenForWebhook("/" + bot.Token)
	go http.ListenAndServeTLS("0.0.0.0:8443", "cert.pem", "key.pem", nil)

	for update := range updates {
		log.Printf("%+v\n", update)
	}
}
```

Nếu cần thiết, bạn có thể tạo một self signed certficate cho việc sử dụng
HTTPS / TLS. Ví dụ ở trên báo cho KingTalk biết đây là
certificate của bạn và bỏ qua việc xác nhận xem chữ ký của certificate có hợp lệ hay không.

    openssl req -x509 -newkey rsa:2048 -keyout key.pem -out cert.pem -days 3560 -subj "//O=Org\CN=Test" -nodes

Hiện tại bạn có thể sử dụng [Let's Encrypt](https://letsencrypt.org) để generate TLS certificate miễn phí.
