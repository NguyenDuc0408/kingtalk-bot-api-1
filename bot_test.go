package ktbotapi_test

import (
	ktbotapi "github.com/nvhai245/kingtalk-bot-api"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"testing"
	"time"
)

const (
	TestToken               = "153667468:AAHlSHlMqSt1f_uFmVRJbm5gntu2HI4WW8I"
	ChatID                  = 76918703
	SupergroupChatID        = -1001120141283
	ReplyToMessageID        = 35
	ExistingPhotoFileID     = "AgADAgADw6cxG4zHKAkr42N7RwEN3IFShCoABHQwXEtVks4EH2wBAAEC"
	ExistingDocumentFileID  = "BQADAgADOQADjMcoCcioX1GrDvp3Ag"
	ExistingAudioFileID     = "BQADAgADRgADjMcoCdXg3lSIN49lAg"
	ExistingVoiceFileID     = "AwADAgADWQADjMcoCeul6r_q52IyAg"
	ExistingVideoFileID     = "BAADAgADZgADjMcoCav432kYe0FRAg"
	ExistingVideoNoteFileID = "DQADAgADdQAD70cQSUK41dLsRMqfAg"
	ExistingStickerFileID   = "BQADAgADcwADjMcoCbdl-6eB--YPAg"
)

func getBot(t *testing.T) (*ktbotapi.BotAPI, error) {
	bot, err := ktbotapi.NewBotAPI(TestToken)
	bot.Debug = true

	if err != nil {
		t.Error(err)
		t.Fail()
	}

	return bot, err
}

func TestNewBotAPI_notoken(t *testing.T) {
	_, err := ktbotapi.NewBotAPI("")

	if err == nil {
		t.Error(err)
		t.Fail()
	}
}

func TestGetUpdates(t *testing.T) {
	bot, _ := getBot(t)

	u := ktbotapi.NewUpdate(0)

	_, err := bot.GetUpdates(u)

	if err != nil {
		t.Error(err)
		t.Fail()
	}
}

func TestSendWithMessage(t *testing.T) {
	bot, _ := getBot(t)

	msg := ktbotapi.NewMessage(ChatID, "A test message from the test library in kingtalk-bot-api")
	msg.ParseMode = "markdown"
	_, err := bot.Send(msg)

	if err != nil {
		t.Error(err)
		t.Fail()
	}
}

func TestSendWithMessageReply(t *testing.T) {
	bot, _ := getBot(t)

	msg := ktbotapi.NewMessage(ChatID, "A test message from the test library in kingtalk-bot-api")
	msg.ReplyToMessageID = ReplyToMessageID
	_, err := bot.Send(msg)

	if err != nil {
		t.Error(err)
		t.Fail()
	}
}

func TestSendWithMessageForward(t *testing.T) {
	bot, _ := getBot(t)

	msg := ktbotapi.NewForward(ChatID, ChatID, ReplyToMessageID)
	_, err := bot.Send(msg)

	if err != nil {
		t.Error(err)
		t.Fail()
	}
}

func TestSendWithNewPhoto(t *testing.T) {
	bot, _ := getBot(t)

	msg := ktbotapi.NewPhotoUpload(ChatID, "examples/image.jpg")
	msg.Caption = "Test"
	_, err := bot.Send(msg)

	if err != nil {
		t.Error(err)
		t.Fail()
	}
}

func TestSendWithNewPhotoWithFileBytes(t *testing.T) {
	bot, _ := getBot(t)

	data, _ := ioutil.ReadFile("examples/image.jpg")
	b := ktbotapi.FileBytes{Name: "image.jpg", Bytes: data}

	msg := ktbotapi.NewPhotoUpload(ChatID, b)
	msg.Caption = "Test"
	_, err := bot.Send(msg)

	if err != nil {
		t.Error(err)
		t.Fail()
	}
}

func TestSendWithNewPhotoWithFileReader(t *testing.T) {
	bot, _ := getBot(t)

	f, _ := os.Open("examples/image.jpg")
	reader := ktbotapi.FileReader{Name: "image.jpg", Reader: f, Size: -1}

	msg := ktbotapi.NewPhotoUpload(ChatID, reader)
	msg.Caption = "Test"
	_, err := bot.Send(msg)

	if err != nil {
		t.Error(err)
		t.Fail()
	}
}

func TestSendWithNewPhotoReply(t *testing.T) {
	bot, _ := getBot(t)

	msg := ktbotapi.NewPhotoUpload(ChatID, "examples/image.jpg")
	msg.ReplyToMessageID = ReplyToMessageID

	_, err := bot.Send(msg)

	if err != nil {
		t.Error(err)
		t.Fail()
	}
}

func TestSendWithExistingPhoto(t *testing.T) {
	bot, _ := getBot(t)

	msg := ktbotapi.NewPhotoShare(ChatID, ExistingPhotoFileID)
	msg.Caption = "Test"
	_, err := bot.Send(msg)

	if err != nil {
		t.Error(err)
		t.Fail()
	}
}

func TestSendWithNewDocument(t *testing.T) {
	bot, _ := getBot(t)

	msg := ktbotapi.NewDocumentUpload(ChatID, "examples/image.jpg")
	_, err := bot.Send(msg)

	if err != nil {
		t.Error(err)
		t.Fail()
	}
}

func TestSendWithExistingDocument(t *testing.T) {
	bot, _ := getBot(t)

	msg := ktbotapi.NewDocumentShare(ChatID, ExistingDocumentFileID)
	_, err := bot.Send(msg)

	if err != nil {
		t.Error(err)
		t.Fail()
	}
}

func TestSendWithNewAudio(t *testing.T) {
	bot, _ := getBot(t)

	msg := ktbotapi.NewAudioUpload(ChatID, "examples/audio.mp3")
	msg.Title = "TEST"
	msg.Duration = 10
	msg.Performer = "TEST"
	msg.MimeType = "audio/mpeg"
	msg.FileSize = 688
	_, err := bot.Send(msg)

	if err != nil {
		t.Error(err)
		t.Fail()
	}
}

func TestSendWithExistingAudio(t *testing.T) {
	bot, _ := getBot(t)

	msg := ktbotapi.NewAudioShare(ChatID, ExistingAudioFileID)
	msg.Title = "TEST"
	msg.Duration = 10
	msg.Performer = "TEST"

	_, err := bot.Send(msg)

	if err != nil {
		t.Error(err)
		t.Fail()
	}
}

func TestSendWithNewVoice(t *testing.T) {
	bot, _ := getBot(t)

	msg := ktbotapi.NewVoiceUpload(ChatID, "examples/voice.ogg")
	msg.Duration = 10
	_, err := bot.Send(msg)

	if err != nil {
		t.Error(err)
		t.Fail()
	}
}

func TestSendWithExistingVoice(t *testing.T) {
	bot, _ := getBot(t)

	msg := ktbotapi.NewVoiceShare(ChatID, ExistingVoiceFileID)
	msg.Duration = 10
	_, err := bot.Send(msg)

	if err != nil {
		t.Error(err)
		t.Fail()
	}
}

func TestSendWithContact(t *testing.T) {
	bot, _ := getBot(t)

	contact := ktbotapi.NewContact(ChatID, "5551234567", "Test")

	if _, err := bot.Send(contact); err != nil {
		t.Error(err)
		t.Fail()
	}
}

func TestSendWithLocation(t *testing.T) {
	bot, _ := getBot(t)

	_, err := bot.Send(ktbotapi.NewLocation(ChatID, 40, 40))

	if err != nil {
		t.Error(err)
		t.Fail()
	}
}

func TestSendWithVenue(t *testing.T) {
	bot, _ := getBot(t)

	venue := ktbotapi.NewVenue(ChatID, "A Test Location", "123 Test Street", 40, 40)

	if _, err := bot.Send(venue); err != nil {
		t.Error(err)
		t.Fail()
	}
}

func TestSendWithNewVideo(t *testing.T) {
	bot, _ := getBot(t)

	msg := ktbotapi.NewVideoUpload(ChatID, "examples/video.mp4")
	msg.Duration = 10
	msg.Caption = "TEST"

	_, err := bot.Send(msg)

	if err != nil {
		t.Error(err)
		t.Fail()
	}
}

func TestSendWithExistingVideo(t *testing.T) {
	bot, _ := getBot(t)

	msg := ktbotapi.NewVideoShare(ChatID, ExistingVideoFileID)
	msg.Duration = 10
	msg.Caption = "TEST"

	_, err := bot.Send(msg)

	if err != nil {
		t.Error(err)
		t.Fail()
	}
}

func TestSendWithNewVideoNote(t *testing.T) {
	bot, _ := getBot(t)

	msg := ktbotapi.NewVideoNoteUpload(ChatID, 240, "examples/videonote.mp4")
	msg.Duration = 10

	_, err := bot.Send(msg)

	if err != nil {
		t.Error(err)
		t.Fail()
	}
}

func TestSendWithExistingVideoNote(t *testing.T) {
	bot, _ := getBot(t)

	msg := ktbotapi.NewVideoNoteShare(ChatID, 240, ExistingVideoNoteFileID)
	msg.Duration = 10

	_, err := bot.Send(msg)

	if err != nil {
		t.Error(err)
		t.Fail()
	}
}

func TestSendWithNewSticker(t *testing.T) {
	bot, _ := getBot(t)

	msg := ktbotapi.NewStickerUpload(ChatID, "examples/image.jpg")

	_, err := bot.Send(msg)

	if err != nil {
		t.Error(err)
		t.Fail()
	}
}

func TestSendWithExistingSticker(t *testing.T) {
	bot, _ := getBot(t)

	msg := ktbotapi.NewStickerShare(ChatID, ExistingStickerFileID)

	_, err := bot.Send(msg)

	if err != nil {
		t.Error(err)
		t.Fail()
	}
}

func TestSendWithNewStickerAndKeyboardHide(t *testing.T) {
	bot, _ := getBot(t)

	msg := ktbotapi.NewStickerUpload(ChatID, "examples/image.jpg")
	msg.ReplyMarkup = ktbotapi.ReplyKeyboardRemove{
		RemoveKeyboard: true,
		Selective:      false,
	}
	_, err := bot.Send(msg)

	if err != nil {
		t.Error(err)
		t.Fail()
	}
}

func TestSendWithExistingStickerAndKeyboardHide(t *testing.T) {
	bot, _ := getBot(t)

	msg := ktbotapi.NewStickerShare(ChatID, ExistingStickerFileID)
	msg.ReplyMarkup = ktbotapi.ReplyKeyboardRemove{
		RemoveKeyboard: true,
		Selective:      false,
	}

	_, err := bot.Send(msg)

	if err != nil {
		t.Error(err)
		t.Fail()
	}
}

func TestGetFile(t *testing.T) {
	bot, _ := getBot(t)

	file := ktbotapi.FileConfig{FileID: ExistingPhotoFileID}

	_, err := bot.GetFile(file)

	if err != nil {
		t.Error(err)
		t.Fail()
	}
}

func TestSendChatConfig(t *testing.T) {
	bot, _ := getBot(t)

	_, err := bot.Send(ktbotapi.NewChatAction(ChatID, ktbotapi.ChatTyping))

	if err != nil {
		t.Error(err)
		t.Fail()
	}
}

func TestSendEditMessage(t *testing.T) {
	bot, _ := getBot(t)

	msg, err := bot.Send(ktbotapi.NewMessage(ChatID, "Testing editing."))
	if err != nil {
		t.Error(err)
		t.Fail()
	}

	edit := ktbotapi.EditMessageTextConfig{
		BaseEdit: ktbotapi.BaseEdit{
			ChatID:    ChatID,
			MessageID: msg.MessageID,
		},
		Text: "Updated text.",
	}

	_, err = bot.Send(edit)
	if err != nil {
		t.Error(err)
		t.Fail()
	}
}

func TestGetUserProfilePhotos(t *testing.T) {
	bot, _ := getBot(t)

	_, err := bot.GetUserProfilePhotos(ktbotapi.NewUserProfilePhotos(ChatID))
	if err != nil {
		t.Error(err)
		t.Fail()
	}
}

func TestSetWebhookWithCert(t *testing.T) {
	bot, _ := getBot(t)

	time.Sleep(time.Second * 2)

	bot.RemoveWebhook()

	wh := ktbotapi.NewWebhookWithCert("https://example.com/ktbotapi-test/"+bot.Token, "examples/cert.pem")
	_, err := bot.SetWebhook(wh)
	if err != nil {
		t.Error(err)
		t.Fail()
	}
	_, err = bot.GetWebhookInfo()
	if err != nil {
		t.Error(err)
	}
	bot.RemoveWebhook()
}

func TestSetWebhookWithoutCert(t *testing.T) {
	bot, _ := getBot(t)

	time.Sleep(time.Second * 2)

	bot.RemoveWebhook()

	wh := ktbotapi.NewWebhook("https://example.com/ktbotapi-test/" + bot.Token)
	_, err := bot.SetWebhook(wh)
	if err != nil {
		t.Error(err)
		t.Fail()
	}
	info, err := bot.GetWebhookInfo()
	if err != nil {
		t.Error(err)
	}
	if info.LastErrorDate != 0 {
		t.Errorf("[KingTalk callback failed]%s", info.LastErrorMessage)
	}
	bot.RemoveWebhook()
}

func TestUpdatesChan(t *testing.T) {
	bot, _ := getBot(t)

	var ucfg ktbotapi.UpdateConfig = ktbotapi.NewUpdate(0)
	ucfg.Timeout = 60
	_, err := bot.GetUpdatesChan(ucfg)

	if err != nil {
		t.Error(err)
		t.Fail()
	}
}

func TestSendWithMediaGroup(t *testing.T) {
	bot, _ := getBot(t)

	cfg := ktbotapi.NewMediaGroup(ChatID, []interface{}{
		ktbotapi.NewInputMediaPhoto("https://i.imgur.com/unQLJIb.jpg"),
		ktbotapi.NewInputMediaPhoto("https://i.imgur.com/J5qweNZ.jpg"),
		ktbotapi.NewInputMediaVideo("https://i.imgur.com/F6RmI24.mp4"),
	})
	_, err := bot.Send(cfg)
	if err != nil {
		t.Error(err)
	}
}

func ExampleNewBotAPI() {
	bot, err := ktbotapi.NewBotAPI("MyAwesomeBotToken")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := ktbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	// Optional: wait for updates and clear them if you don't want to handle
	// a large backlog of old messages
	time.Sleep(time.Millisecond * 500)
	updates.Clear()

	for update := range updates {
		if update.Message == nil {
			continue
		}

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		msg := ktbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
		msg.ReplyToMessageID = update.Message.MessageID

		bot.Send(msg)
	}
}

func ExampleNewWebhook() {
	bot, err := ktbotapi.NewBotAPI("MyAwesomeBotToken")
	if err != nil {
		log.Fatal(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	_, err = bot.SetWebhook(ktbotapi.NewWebhookWithCert("https://www.google.com:8443/"+bot.Token, "cert.pem"))
	if err != nil {
		log.Fatal(err)
	}
	info, err := bot.GetWebhookInfo()
	if err != nil {
		log.Fatal(err)
	}
	if info.LastErrorDate != 0 {
		log.Printf("[KingTalk callback failed]%s", info.LastErrorMessage)
	}
	updates := bot.ListenForWebhook("/" + bot.Token)
	go http.ListenAndServeTLS("0.0.0.0:8443", "cert.pem", "key.pem", nil)

	for update := range updates {
		log.Printf("%+v\n", update)
	}
}

func ExampleAnswerInlineQuery() {
	bot, err := ktbotapi.NewBotAPI("MyAwesomeBotToken") // create new bot
	if err != nil {
		log.Panic(err)
	}

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := ktbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.InlineQuery == nil { // if no inline query, ignore it
			continue
		}

		article := ktbotapi.NewInlineQueryResultArticle(update.InlineQuery.ID, "Echo", update.InlineQuery.Query)
		article.Description = update.InlineQuery.Query

		inlineConf := ktbotapi.InlineConfig{
			InlineQueryID: update.InlineQuery.ID,
			IsPersonal:    true,
			CacheTime:     0,
			Results:       []interface{}{article},
		}

		if _, err := bot.AnswerInlineQuery(inlineConf); err != nil {
			log.Println(err)
		}
	}
}

func TestDeleteMessage(t *testing.T) {
	bot, _ := getBot(t)

	msg := ktbotapi.NewMessage(ChatID, "A test message from the test library in telegram-bot-api")
	msg.ParseMode = "markdown"
	message, _ := bot.Send(msg)

	deleteMessageConfig := ktbotapi.DeleteMessageConfig{
		ChatID:    message.Chat.ID,
		MessageID: message.MessageID,
	}
	_, err := bot.DeleteMessage(deleteMessageConfig)

	if err != nil {
		t.Error(err)
		t.Fail()
	}
}

func TestPinChatMessage(t *testing.T) {
	bot, _ := getBot(t)

	msg := ktbotapi.NewMessage(SupergroupChatID, "A test message from the test library in kingtalk-bot-api")
	msg.ParseMode = "markdown"
	message, _ := bot.Send(msg)

	pinChatMessageConfig := ktbotapi.PinChatMessageConfig{
		ChatID:              message.Chat.ID,
		MessageID:           message.MessageID,
		DisableNotification: false,
	}
	_, err := bot.PinChatMessage(pinChatMessageConfig)

	if err != nil {
		t.Error(err)
		t.Fail()
	}
}

func TestUnpinChatMessage(t *testing.T) {
	bot, _ := getBot(t)

	msg := ktbotapi.NewMessage(SupergroupChatID, "A test message from the test library in kingtalk-bot-api")
	msg.ParseMode = "markdown"
	message, _ := bot.Send(msg)

	// We need pin message to unpin something
	pinChatMessageConfig := ktbotapi.PinChatMessageConfig{
		ChatID:              message.Chat.ID,
		MessageID:           message.MessageID,
		DisableNotification: false,
	}
	_, err := bot.PinChatMessage(pinChatMessageConfig)

	unpinChatMessageConfig := ktbotapi.UnpinChatMessageConfig{
		ChatID: message.Chat.ID,
	}
	_, err = bot.UnpinChatMessage(unpinChatMessageConfig)

	if err != nil {
		t.Error(err)
		t.Fail()
	}
}
