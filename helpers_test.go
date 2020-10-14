package ktbotapi_test

import (
	"github.com/nvhai245/kingtalk-bot-api"
	"testing"
)

func TestNewInlineQueryResultArticle(t *testing.T) {
	result := ktbotapi.NewInlineQueryResultArticle("id", "title", "message")

	if result.Type != "article" ||
		result.ID != "id" ||
		result.Title != "title" ||
		result.InputMessageContent.(ktbotapi.InputTextMessageContent).Text != "message" {
		t.Fail()
	}
}

func TestNewInlineQueryResultArticleMarkdown(t *testing.T) {
	result := ktbotapi.NewInlineQueryResultArticleMarkdown("id", "title", "*message*")

	if result.Type != "article" ||
		result.ID != "id" ||
		result.Title != "title" ||
		result.InputMessageContent.(ktbotapi.InputTextMessageContent).Text != "*message*" ||
		result.InputMessageContent.(ktbotapi.InputTextMessageContent).ParseMode != "Markdown" {
		t.Fail()
	}
}

func TestNewInlineQueryResultArticleHTML(t *testing.T) {
	result := ktbotapi.NewInlineQueryResultArticleHTML("id", "title", "<b>message</b>")

	if result.Type != "article" ||
		result.ID != "id" ||
		result.Title != "title" ||
		result.InputMessageContent.(ktbotapi.InputTextMessageContent).Text != "<b>message</b>" ||
		result.InputMessageContent.(ktbotapi.InputTextMessageContent).ParseMode != "HTML" {
		t.Fail()
	}
}

func TestNewInlineQueryResultGIF(t *testing.T) {
	result := ktbotapi.NewInlineQueryResultGIF("id", "google.com")

	if result.Type != "gif" ||
		result.ID != "id" ||
		result.URL != "google.com" {
		t.Fail()
	}
}

func TestNewInlineQueryResultMPEG4GIF(t *testing.T) {
	result := ktbotapi.NewInlineQueryResultMPEG4GIF("id", "google.com")

	if result.Type != "mpeg4_gif" ||
		result.ID != "id" ||
		result.URL != "google.com" {
		t.Fail()
	}
}

func TestNewInlineQueryResultPhoto(t *testing.T) {
	result := ktbotapi.NewInlineQueryResultPhoto("id", "google.com")

	if result.Type != "photo" ||
		result.ID != "id" ||
		result.URL != "google.com" {
		t.Fail()
	}
}

func TestNewInlineQueryResultPhotoWithThumb(t *testing.T) {
	result := ktbotapi.NewInlineQueryResultPhotoWithThumb("id", "google.com", "thumb.com")

	if result.Type != "photo" ||
		result.ID != "id" ||
		result.URL != "google.com" ||
		result.ThumbURL != "thumb.com" {
		t.Fail()
	}
}

func TestNewInlineQueryResultVideo(t *testing.T) {
	result := ktbotapi.NewInlineQueryResultVideo("id", "google.com")

	if result.Type != "video" ||
		result.ID != "id" ||
		result.URL != "google.com" {
		t.Fail()
	}
}

func TestNewInlineQueryResultAudio(t *testing.T) {
	result := ktbotapi.NewInlineQueryResultAudio("id", "google.com", "title")

	if result.Type != "audio" ||
		result.ID != "id" ||
		result.URL != "google.com" ||
		result.Title != "title" {
		t.Fail()
	}
}

func TestNewInlineQueryResultVoice(t *testing.T) {
	result := ktbotapi.NewInlineQueryResultVoice("id", "google.com", "title")

	if result.Type != "voice" ||
		result.ID != "id" ||
		result.URL != "google.com" ||
		result.Title != "title" {
		t.Fail()
	}
}

func TestNewInlineQueryResultDocument(t *testing.T) {
	result := ktbotapi.NewInlineQueryResultDocument("id", "google.com", "title", "mime/type")

	if result.Type != "document" ||
		result.ID != "id" ||
		result.URL != "google.com" ||
		result.Title != "title" ||
		result.MimeType != "mime/type" {
		t.Fail()
	}
}

func TestNewInlineQueryResultLocation(t *testing.T) {
	result := ktbotapi.NewInlineQueryResultLocation("id", "name", 40, 50)

	if result.Type != "location" ||
		result.ID != "id" ||
		result.Title != "name" ||
		result.Latitude != 40 ||
		result.Longitude != 50 {
		t.Fail()
	}
}

func TestNewEditMessageText(t *testing.T) {
	edit := ktbotapi.NewEditMessageText(ChatID, ReplyToMessageID, "new text")

	if edit.Text != "new text" ||
		edit.BaseEdit.ChatID != ChatID ||
		edit.BaseEdit.MessageID != ReplyToMessageID {
		t.Fail()
	}
}

func TestNewEditMessageCaption(t *testing.T) {
	edit := ktbotapi.NewEditMessageCaption(ChatID, ReplyToMessageID, "new caption")

	if edit.Caption != "new caption" ||
		edit.BaseEdit.ChatID != ChatID ||
		edit.BaseEdit.MessageID != ReplyToMessageID {
		t.Fail()
	}
}

func TestNewEditMessageReplyMarkup(t *testing.T) {
	markup := ktbotapi.InlineKeyboardMarkup{
		InlineKeyboard: [][]ktbotapi.InlineKeyboardButton{
			[]ktbotapi.InlineKeyboardButton{
				ktbotapi.InlineKeyboardButton{Text: "test"},
			},
		},
	}

	edit := ktbotapi.NewEditMessageReplyMarkup(ChatID, ReplyToMessageID, markup)

	if edit.ReplyMarkup.InlineKeyboard[0][0].Text != "test" ||
		edit.BaseEdit.ChatID != ChatID ||
		edit.BaseEdit.MessageID != ReplyToMessageID {
		t.Fail()
	}

}
