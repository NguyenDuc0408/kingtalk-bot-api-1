package ktbotapi

import (
	"net/url"
)

// NewMessage creates a new Message.
//
// chatID is where to send it, text is the message text.
func NewMessage(chatID int64, text string) MessageConfig {
	return MessageConfig{
		BaseChat: BaseChat{
			ChatID:           chatID,
			ReplyToMessageID: 0,
		},
		Text:                  text,
		DisableWebPagePreview: false,
	}
}

// NewMessage creates a new Message.
//
// peerID is the id of the user to send it to, text is the message text.
//
// This method only available for bots with privileges.
func NewMessageWithPeer(peerID int32, text string) MessageConfig {
	return MessageConfig{
		BaseChat: BaseChat{
			PeerID:           peerID,
			ReplyToMessageID: 0,
		},
		Text:                  text,
		DisableWebPagePreview: false,
	}
}

// NewDeleteMessage creates a request to delete a message.
func NewDeleteMessage(chatID int64, messageID int) DeleteMessageConfig {
	return DeleteMessageConfig{
		ChatID:    chatID,
		MessageID: messageID,
	}
}

// NewMessageToChannel creates a new Message that is sent to a channel
// by username.
//
// username is the username of the channel, text is the message text,
// and the username should be in the form of `@username`.
func NewMessageToChannel(username string, text string) MessageConfig {
	return MessageConfig{
		BaseChat: BaseChat{
			ChannelUsername: username,
		},
		Text: text,
	}
}

// NewForward creates a new forward.
//
// chatID is where to send it, fromChatID is the source chat,
// and messageID is the ID of the original message.
func NewForward(chatID int64, fromChatID int64, messageID int) ForwardConfig {
	return ForwardConfig{
		BaseChat:   BaseChat{ChatID: chatID},
		FromChatID: fromChatID,
		MessageID:  messageID,
	}
}

// NewPhotoUpload creates a new photo uploader.
//
// chatID is where to send it, file is a string path to the file,
// FileReader, or FileBytes.
//
// Note that you must send animated GIFs as a document.
func NewPhotoUpload(chatID int64, file interface{}) PhotoConfig {
	return PhotoConfig{
		BaseFile: BaseFile{
			BaseChat:    BaseChat{ChatID: chatID},
			File:        file,
			UseExisting: false,
		},
	}
}

// NewPhotoUploadWithPeer creates a new photo uploader.
//
// peerID is the id of the user to send it to, file is a string path to the file,
// FileReader, or FileBytes.
//
// Note that you must send animated GIFs as a document.
// This method only work for bots with privileges.
func NewPhotoUploadWithPeer(peerID int32, file interface{}) PhotoConfig {
	return PhotoConfig{
		BaseFile: BaseFile{
			BaseChat:    BaseChat{PeerID: peerID},
			File:        file,
			UseExisting: false,
		},
	}
}

// NewPhotoShare shares an existing photo.
// You may use this to reshare an existing photo without reuploading it.
//
// chatID is where to send it, fileID is the ID of the file
// already uploaded.
func NewPhotoShare(chatID int64, fileID string) PhotoConfig {
	return PhotoConfig{
		BaseFile: BaseFile{
			BaseChat:    BaseChat{ChatID: chatID},
			FileID:      fileID,
			UseExisting: true,
		},
	}
}

// NewPhotoShareWithPeer shares an existing photo to an user.
// You may use this to reshare an existing photo without reuploading it.
//
// peerID is the id of the user to send it to, fileID is the ID of the file
// already uploaded.
// This method only work for bots with privileges.
func NewPhotoShareWithPeer(peerID int32, fileID string) PhotoConfig {
	return PhotoConfig{
		BaseFile: BaseFile{
			BaseChat:    BaseChat{PeerID: peerID},
			FileID:      fileID,
			UseExisting: true,
		},
	}
}

// NewAudioUpload creates a new audio uploader.
//
// chatID is where to send it, file is a string path to the file,
// FileReader, or FileBytes.
func NewAudioUpload(chatID int64, file interface{}) AudioConfig {
	return AudioConfig{
		BaseFile: BaseFile{
			BaseChat:    BaseChat{ChatID: chatID},
			File:        file,
			UseExisting: false,
		},
	}
}

// NewAudioUploadWithPeer creates a new audio uploader.
//
// peerID is the id of the user to send it to, file is a string path to the file,
// FileReader, or FileBytes.
//
// This method only available for bots with privileges.
func NewAudioUploadWithPeer(peerID int32, file interface{}) AudioConfig {
	return AudioConfig{
		BaseFile: BaseFile{
			BaseChat:    BaseChat{PeerID: peerID},
			File:        file,
			UseExisting: false,
		},
	}
}

// NewAudioShare shares an existing audio file.
// You may use this to reshare an existing audio file without
// reuploading it.
//
// chatID is where to send it, fileID is the ID of the audio
// already uploaded.
func NewAudioShare(chatID int64, fileID string) AudioConfig {
	return AudioConfig{
		BaseFile: BaseFile{
			BaseChat:    BaseChat{ChatID: chatID},
			FileID:      fileID,
			UseExisting: true,
		},
	}
}

// NewAudioShareWithPeer shares an existing audio file.
// You may use this to reshare an existing audio file without
// reuploading it.
//
// peerID is the id of the user to send it to, fileID is the ID of the audio
// already uploaded.
//
// This method only available for bots with privileges.
func NewAudioShareWithPeer(peerID int32, fileID string) AudioConfig {
	return AudioConfig{
		BaseFile: BaseFile{
			BaseChat:    BaseChat{PeerID: peerID},
			FileID:      fileID,
			UseExisting: true,
		},
	}
}

// NewDocumentUpload creates a new document uploader.
//
// chatID is where to send it, file is a string path to the file,
// FileReader, or FileBytes.
func NewDocumentUpload(chatID int64, file interface{}) DocumentConfig {
	return DocumentConfig{
		BaseFile: BaseFile{
			BaseChat:    BaseChat{ChatID: chatID},
			File:        file,
			UseExisting: false,
		},
	}
}

// NewDocumentUploadWithPeer creates a new document uploader.
//
// peerID is the id of the user to send it to, file is a string path to the file,
// FileReader, or FileBytes.
//
// This method only available for bots with privileges.
func NewDocumentUploadWithPeer(peerID int32, file interface{}) DocumentConfig {
	return DocumentConfig{
		BaseFile: BaseFile{
			BaseChat:    BaseChat{PeerID: peerID},
			File:        file,
			UseExisting: false,
		},
	}
}

// NewDocumentShare shares an existing document.
// You may use this to reshare an existing document without
// reuploading it.
//
// chatID is where to send it, fileID is the ID of the document
// already uploaded.
func NewDocumentShare(chatID int64, fileID string) DocumentConfig {
	return DocumentConfig{
		BaseFile: BaseFile{
			BaseChat:    BaseChat{ChatID: chatID},
			FileID:      fileID,
			UseExisting: true,
		},
	}
}

// NewDocumentShareWithPeer shares an existing document.
// You may use this to reshare an existing document without
// reuploading it.
//
// peerID is the id of the user to send it to, fileID is the ID of the document
// already uploaded.
//
// This method only available for bots with privileges.
func NewDocumentShareWithPeer(peerID int32, fileID string) DocumentConfig {
	return DocumentConfig{
		BaseFile: BaseFile{
			BaseChat:    BaseChat{PeerID: peerID},
			FileID:      fileID,
			UseExisting: true,
		},
	}
}

// NewStickerUpload creates a new sticker uploader.
//
// chatID is where to send it, file is a string path to the file,
// FileReader, or FileBytes.
func NewStickerUpload(chatID int64, file interface{}) StickerConfig {
	return StickerConfig{
		BaseFile: BaseFile{
			BaseChat:    BaseChat{ChatID: chatID},
			File:        file,
			UseExisting: false,
		},
	}
}

// NewStickerUploadWithPeer creates a new sticker uploader.
//
// peerID is the id of the user to send it to, file is a string path to the file,
// FileReader, or FileBytes.
//
// This method only available for bots with privileges.
func NewStickerUploadWithPeer(peerID int32, file interface{}) StickerConfig {
	return StickerConfig{
		BaseFile: BaseFile{
			BaseChat:    BaseChat{PeerID: peerID},
			File:        file,
			UseExisting: false,
		},
	}
}

// NewStickerShare shares an existing sticker.
// You may use this to reshare an existing sticker without
// reuploading it.
//
// chatID is where to send it, fileID is the ID of the sticker
// already uploaded.
func NewStickerShare(chatID int64, fileID string) StickerConfig {
	return StickerConfig{
		BaseFile: BaseFile{
			BaseChat:    BaseChat{ChatID: chatID},
			FileID:      fileID,
			UseExisting: true,
		},
	}
}

// NewStickerShareWithPeer shares an existing sticker.
// You may use this to reshare an existing sticker without
// reuploading it.
//
// peerID is the id of the user to send it to, fileID is the ID of the sticker
// already uploaded.
//
// This method only available for bots with privileges.
func NewStickerShareWithPeer(peerID int32, fileID string) StickerConfig {
	return StickerConfig{
		BaseFile: BaseFile{
			BaseChat:    BaseChat{PeerID: peerID},
			FileID:      fileID,
			UseExisting: true,
		},
	}
}

// NewVideoUpload creates a new video uploader.
//
// peerID is the id of the user to send it to, file is a string path to the file,
// FileReader, or FileBytes.
func NewVideoUpload(chatID int64, file interface{}) VideoConfig {
	return VideoConfig{
		BaseFile: BaseFile{
			BaseChat:    BaseChat{ChatID: chatID},
			File:        file,
			UseExisting: false,
		},
	}
}

// NewVideoUploadWithPeer creates a new video uploader.
//
// peerID is the id of the user to send it to, file is a string path to the file,
// FileReader, or FileBytes.
//
// This method only available for bots with privileges.
func NewVideoUploadWithPeer(peerID int32, file interface{}) VideoConfig {
	return VideoConfig{
		BaseFile: BaseFile{
			BaseChat:    BaseChat{PeerID: peerID},
			File:        file,
			UseExisting: false,
		},
	}
}

// NewVideoShare shares an existing video.
// You may use this to reshare an existing video without reuploading it.
//
// chatID is where to send it, fileID is the ID of the video
// already uploaded.
func NewVideoShare(chatID int64, fileID string) VideoConfig {
	return VideoConfig{
		BaseFile: BaseFile{
			BaseChat:    BaseChat{ChatID: chatID},
			FileID:      fileID,
			UseExisting: true,
		},
	}
}

// NewVideoShareWithPeer shares an existing video.
// You may use this to reshare an existing video without reuploading it.
//
// peerID is the id of the user to send it to, fileID is the ID of the video
// already uploaded.
//
// This method only available for bots with privileges.
func NewVideoShareWithPeer(peerID int32, fileID string) VideoConfig {
	return VideoConfig{
		BaseFile: BaseFile{
			BaseChat:    BaseChat{PeerID: peerID},
			FileID:      fileID,
			UseExisting: true,
		},
	}
}

// NewAnimationUpload creates a new animation uploader.
//
// chatID is where to send it, file is a string path to the file,
// FileReader, or FileBytes.
func NewAnimationUpload(chatID int64, file interface{}) AnimationConfig {
	return AnimationConfig{
		BaseFile: BaseFile{
			BaseChat:    BaseChat{ChatID: chatID},
			File:        file,
			UseExisting: false,
		},
	}
}

// NewAnimationUploadWithPeer creates a new animation uploader.
//
// peerID is the id of the user to send it to, file is a string path to the file,
// FileReader, or FileBytes.
//
// This method only available for bots with privileges.
func NewAnimationUploadWithPeer(peerID int32, file interface{}) AnimationConfig {
	return AnimationConfig{
		BaseFile: BaseFile{
			BaseChat:    BaseChat{PeerID: peerID},
			File:        file,
			UseExisting: false,
		},
	}
}

// NewAnimationShare shares an existing animation.
// You may use this to reshare an existing animation without reuploading it.
//
// chatID is where to send it, fileID is the ID of the animation
// already uploaded.
func NewAnimationShare(chatID int64, fileID string) AnimationConfig {
	return AnimationConfig{
		BaseFile: BaseFile{
			BaseChat:    BaseChat{ChatID: chatID},
			FileID:      fileID,
			UseExisting: true,
		},
	}
}

// NewAnimationShareWithPeer shares an existing animation.
// You may use this to reshare an existing animation without reuploading it.
//
// peerID is the id of the user to send it to, fileID is the ID of the animation
// already uploaded.
//
// This method only available for bots with privileges.
func NewAnimationShareWithPeer(peerID int32, fileID string) AnimationConfig {
	return AnimationConfig{
		BaseFile: BaseFile{
			BaseChat:    BaseChat{PeerID: peerID},
			FileID:      fileID,
			UseExisting: true,
		},
	}
}

// NewVideoNoteUpload creates a new video note uploader.
//
// chatID is where to send it, file is a string path to the file,
// FileReader, or FileBytes.
func NewVideoNoteUpload(chatID int64, length int, file interface{}) VideoNoteConfig {
	return VideoNoteConfig{
		BaseFile: BaseFile{
			BaseChat:    BaseChat{ChatID: chatID},
			File:        file,
			UseExisting: false,
		},
		Length: length,
	}
}

// NewVideoNoteUploadWithPeer creates a new video note uploader.
//
// peerID is the id of the user to send it to, file is a string path to the file,
// FileReader, or FileBytes.
//
// This method only available for bots with privileges.
func NewVideoNoteUploadWithPeer(peerID int32, length int, file interface{}) VideoNoteConfig {
	return VideoNoteConfig{
		BaseFile: BaseFile{
			BaseChat:    BaseChat{PeerID: peerID},
			File:        file,
			UseExisting: false,
		},
		Length: length,
	}
}

// NewVideoNoteShare shares an existing video.
// You may use this to reshare an existing video without reuploading it.
//
// chatID is where to send it, fileID is the ID of the video
// already uploaded.
func NewVideoNoteShare(chatID int64, length int, fileID string) VideoNoteConfig {
	return VideoNoteConfig{
		BaseFile: BaseFile{
			BaseChat:    BaseChat{ChatID: chatID},
			FileID:      fileID,
			UseExisting: true,
		},
		Length: length,
	}
}

// NewVideoNoteShareWithPeer shares an existing video.
// You may use this to reshare an existing video without reuploading it.
//
// peerID is the id of the user to send it to, fileID is the ID of the video
// already uploaded.
//
// This method only available for bots with privileges.
func NewVideoNoteShareWithPeer(peerID int32, length int, fileID string) VideoNoteConfig {
	return VideoNoteConfig{
		BaseFile: BaseFile{
			BaseChat:    BaseChat{PeerID: peerID},
			FileID:      fileID,
			UseExisting: true,
		},
		Length: length,
	}
}

// NewVoiceUpload creates a new voice uploader.
//
// chatID is where to send it, file is a string path to the file,
// FileReader, or FileBytes.
func NewVoiceUpload(chatID int64, file interface{}) VoiceConfig {
	return VoiceConfig{
		BaseFile: BaseFile{
			BaseChat:    BaseChat{ChatID: chatID},
			File:        file,
			UseExisting: false,
		},
	}
}

// NewVoiceUploadWithPeer creates a new voice uploader.
//
// peerID is the id of the user to send it to, file is a string path to the file,
// FileReader, or FileBytes.
//
// This method only available for bots with privileges.
func NewVoiceUploadWithPeer(peerID int32, file interface{}) VoiceConfig {
	return VoiceConfig{
		BaseFile: BaseFile{
			BaseChat:    BaseChat{PeerID: peerID},
			File:        file,
			UseExisting: false,
		},
	}
}

// NewVoiceShare shares an existing voice.
// You may use this to reshare an existing voice without reuploading it.
//
// chatID is where to send it, fileID is the ID of the video
// already uploaded.
func NewVoiceShare(chatID int64, fileID string) VoiceConfig {
	return VoiceConfig{
		BaseFile: BaseFile{
			BaseChat:    BaseChat{ChatID: chatID},
			FileID:      fileID,
			UseExisting: true,
		},
	}
}

// NewVoiceShareWithPeer shares an existing voice.
// You may use this to reshare an existing voice without reuploading it.
//
// peerID is the id of the user to send it to, fileID is the ID of the video
// already uploaded.
//
// This method only available for bots with privileges.
func NewVoiceShareWithPeer(peerID int32, fileID string) VoiceConfig {
	return VoiceConfig{
		BaseFile: BaseFile{
			BaseChat:    BaseChat{PeerID: peerID},
			FileID:      fileID,
			UseExisting: true,
		},
	}
}

// NewMediaGroup creates a new media group. Files should be an array of
// two to ten InputMediaPhoto or InputMediaVideo.
func NewMediaGroup(chatID int64, files []interface{}) MediaGroupConfig {
	return MediaGroupConfig{
		BaseChat: BaseChat{
			ChatID: chatID,
		},
		InputMedia: files,
	}
}

// NewMediaGroupWithPeer creates a new media group. Files should be an array of
// two to ten InputMediaPhoto or InputMediaVideo.
// peerID is the id of the user to send it to.
//
// This method only available for bots with privileges.
func NewMediaGroupWithPeer(peerID int32, files []interface{}) MediaGroupConfig {
	return MediaGroupConfig{
		BaseChat: BaseChat{
			PeerID: peerID,
		},
		InputMedia: files,
	}
}

// NewInputMediaPhoto creates a new InputMediaPhoto.
func NewInputMediaPhoto(media string) InputMediaPhoto {
	return InputMediaPhoto{
		Type:  "photo",
		Media: media,
	}
}

// NewInputMediaVideo creates a new InputMediaVideo.
func NewInputMediaVideo(media string) InputMediaVideo {
	return InputMediaVideo{
		Type:  "video",
		Media: media,
	}
}

// NewContact allows you to send a shared contact.
func NewContact(chatID int64, phoneNumber, firstName string) ContactConfig {
	return ContactConfig{
		BaseChat: BaseChat{
			ChatID: chatID,
		},
		PhoneNumber: phoneNumber,
		FirstName:   firstName,
	}
}

// NewLocation shares your location.
//
// chatID is where to send it, latitude and longitude are coordinates.
func NewLocation(chatID int64, latitude float64, longitude float64) LocationConfig {
	return LocationConfig{
		BaseChat: BaseChat{
			ChatID: chatID,
		},
		Latitude:  latitude,
		Longitude: longitude,
	}
}

// NewVenue allows you to send a venue and its location.
func NewVenue(chatID int64, title, address string, latitude, longitude float64) VenueConfig {
	return VenueConfig{
		BaseChat: BaseChat{
			ChatID: chatID,
		},
		Title:     title,
		Address:   address,
		Latitude:  latitude,
		Longitude: longitude,
	}
}

// NewChatAction sets a chat action.
// Actions last for 5 seconds, or until your next action.
//
// chatID is where to send it, action should be set via Chat constants.
func NewChatAction(chatID int64, action string) ChatActionConfig {
	return ChatActionConfig{
		BaseChat: BaseChat{ChatID: chatID},
		Action:   action,
	}
}

// NewUserProfilePhotos gets user profile photos.
//
// userID is the ID of the user you wish to get profile photos from.
func NewUserProfilePhotos(userID int) UserProfilePhotosConfig {
	return UserProfilePhotosConfig{
		UserID: userID,
		Offset: 0,
		Limit:  0,
	}
}

// NewUpdate gets updates since the last Offset.
//
// offset is the last Update ID to include.
// You likely want to set this to the last Update ID plus 1.
func NewUpdate(offset int) UpdateConfig {
	return UpdateConfig{
		Offset:  offset,
		Limit:   0,
		Timeout: 0,
	}
}

// NewWebhook creates a new webhook.
//
// link is the url parsable link you wish to get the updates.
func NewWebhook(link string) WebhookConfig {
	u, _ := url.Parse(link)

	return WebhookConfig{
		URL: u,
	}
}

// NewWebhookWithCert creates a new webhook with a certificate.
//
// link is the url you wish to get webhooks,
// file contains a string to a file, FileReader, or FileBytes.
func NewWebhookWithCert(link string, file interface{}) WebhookConfig {
	u, _ := url.Parse(link)

	return WebhookConfig{
		URL:         u,
		Certificate: file,
	}
}

// NewInlineQueryResultArticle creates a new inline query article.
func NewInlineQueryResultArticle(id, title, messageText string) InlineQueryResultArticle {
	return InlineQueryResultArticle{
		Type:  "article",
		ID:    id,
		Title: title,
		InputMessageContent: InputTextMessageContent{
			Text: messageText,
		},
	}
}

// NewInlineQueryResultArticleMarkdown creates a new inline query article with Markdown parsing.
func NewInlineQueryResultArticleMarkdown(id, title, messageText string) InlineQueryResultArticle {
	return InlineQueryResultArticle{
		Type:  "article",
		ID:    id,
		Title: title,
		InputMessageContent: InputTextMessageContent{
			Text:      messageText,
			ParseMode: "Markdown",
		},
	}
}

// NewInlineQueryResultArticleHTML creates a new inline query article with HTML parsing.
func NewInlineQueryResultArticleHTML(id, title, messageText string) InlineQueryResultArticle {
	return InlineQueryResultArticle{
		Type:  "article",
		ID:    id,
		Title: title,
		InputMessageContent: InputTextMessageContent{
			Text:      messageText,
			ParseMode: "HTML",
		},
	}
}

// NewInlineQueryResultGIF creates a new inline query GIF.
func NewInlineQueryResultGIF(id, url string) InlineQueryResultGIF {
	return InlineQueryResultGIF{
		Type: "gif",
		ID:   id,
		URL:  url,
	}
}

// NewInlineQueryResultCachedGIF create a new inline query with cached photo.
func NewInlineQueryResultCachedGIF(id, gifID string) InlineQueryResultCachedGIF {
	return InlineQueryResultCachedGIF{
		Type:  "gif",
		ID:    id,
		GifID: gifID,
	}
}

// NewInlineQueryResultMPEG4GIF creates a new inline query MPEG4 GIF.
func NewInlineQueryResultMPEG4GIF(id, url string) InlineQueryResultMPEG4GIF {
	return InlineQueryResultMPEG4GIF{
		Type: "mpeg4_gif",
		ID:   id,
		URL:  url,
	}
}

// NewInlineQueryResultCachedPhoto create a new inline query with cached photo.
func NewInlineQueryResultCachedMPEG4GIF(id, MPEG4GifID string) InlineQueryResultCachedMpeg4Gif {
	return InlineQueryResultCachedMpeg4Gif{
		Type:   "mpeg4_gif",
		ID:     id,
		MGifID: MPEG4GifID,
	}
}

// NewInlineQueryResultPhoto creates a new inline query photo.
func NewInlineQueryResultPhoto(id, url string) InlineQueryResultPhoto {
	return InlineQueryResultPhoto{
		Type: "photo",
		ID:   id,
		URL:  url,
	}
}

// NewInlineQueryResultPhotoWithThumb creates a new inline query photo.
func NewInlineQueryResultPhotoWithThumb(id, url, thumb string) InlineQueryResultPhoto {
	return InlineQueryResultPhoto{
		Type:     "photo",
		ID:       id,
		URL:      url,
		ThumbURL: thumb,
	}
}

// NewInlineQueryResultCachedPhoto create a new inline query with cached photo.
func NewInlineQueryResultCachedPhoto(id, photoID string) InlineQueryResultCachedPhoto {
	return InlineQueryResultCachedPhoto{
		Type:    "photo",
		ID:      id,
		PhotoID: photoID,
	}
}

// NewInlineQueryResultVideo creates a new inline query video.
func NewInlineQueryResultVideo(id, url string) InlineQueryResultVideo {
	return InlineQueryResultVideo{
		Type: "video",
		ID:   id,
		URL:  url,
	}
}

// NewInlineQueryResultCachedVideo create a new inline query with cached video.
func NewInlineQueryResultCachedVideo(id, videoID, title string) InlineQueryResultCachedVideo {
	return InlineQueryResultCachedVideo{
		Type:    "video",
		ID:      id,
		VideoID: videoID,
		Title:   title,
	}
}

// NewInlineQueryResultAudio creates a new inline query audio.
func NewInlineQueryResultAudio(id, url, title string) InlineQueryResultAudio {
	return InlineQueryResultAudio{
		Type:  "audio",
		ID:    id,
		URL:   url,
		Title: title,
	}
}

// NewInlineQueryResultCachedAudio create a new inline query with cached photo.
func NewInlineQueryResultCachedAudio(id, audioID string) InlineQueryResultCachedAudio {
	return InlineQueryResultCachedAudio{
		Type:    "audio",
		ID:      id,
		AudioID: audioID,
	}
}

// NewInlineQueryResultVoice creates a new inline query voice.
func NewInlineQueryResultVoice(id, url, title string) InlineQueryResultVoice {
	return InlineQueryResultVoice{
		Type:  "voice",
		ID:    id,
		URL:   url,
		Title: title,
	}
}

// NewInlineQueryResultCachedVoice create a new inline query with cached photo.
func NewInlineQueryResultCachedVoice(id, voiceID, title string) InlineQueryResultCachedVoice {
	return InlineQueryResultCachedVoice{
		Type:    "voice",
		ID:      id,
		VoiceID: voiceID,
		Title:   title,
	}
}

// NewInlineQueryResultDocument creates a new inline query document.
func NewInlineQueryResultDocument(id, url, title, mimeType string) InlineQueryResultDocument {
	return InlineQueryResultDocument{
		Type:     "document",
		ID:       id,
		URL:      url,
		Title:    title,
		MimeType: mimeType,
	}
}

// NewInlineQueryResultCachedDocument create a new inline query with cached photo.
func NewInlineQueryResultCachedDocument(id, documentID, title string) InlineQueryResultCachedDocument {
	return InlineQueryResultCachedDocument{
		Type:       "document",
		ID:         id,
		DocumentID: documentID,
		Title:      title,
	}
}

// NewInlineQueryResultLocation creates a new inline query location.
func NewInlineQueryResultLocation(id, title string, latitude, longitude float64) InlineQueryResultLocation {
	return InlineQueryResultLocation{
		Type:      "location",
		ID:        id,
		Title:     title,
		Latitude:  latitude,
		Longitude: longitude,
	}
}

// NewInlineQueryResultVenue creates a new inline query venue.
func NewInlineQueryResultVenue(id, title, address string, latitude, longitude float64) InlineQueryResultVenue {
	return InlineQueryResultVenue{
		Type:      "venue",
		ID:        id,
		Title:     title,
		Address:   address,
		Latitude:  latitude,
		Longitude: longitude,
	}
}

// NewEditMessageText allows you to edit the text of a message.
func NewEditMessageText(chatID int64, messageID int, text string) EditMessageTextConfig {
	return EditMessageTextConfig{
		BaseEdit: BaseEdit{
			ChatID:    chatID,
			MessageID: messageID,
		},
		Text: text,
	}
}

// NewEditMessageCaption allows you to edit the caption of a message.
func NewEditMessageCaption(chatID int64, messageID int, caption string) EditMessageCaptionConfig {
	return EditMessageCaptionConfig{
		BaseEdit: BaseEdit{
			ChatID:    chatID,
			MessageID: messageID,
		},
		Caption: caption,
	}
}

// NewEditMessageReplyMarkup allows you to edit the inline
// keyboard markup.
func NewEditMessageReplyMarkup(chatID int64, messageID int, replyMarkup InlineKeyboardMarkup) EditMessageReplyMarkupConfig {
	return EditMessageReplyMarkupConfig{
		BaseEdit: BaseEdit{
			ChatID:      chatID,
			MessageID:   messageID,
			ReplyMarkup: &replyMarkup,
		},
	}
}

// NewHideKeyboard hides the keyboard, with the option for being selective
// or hiding for everyone.
func NewHideKeyboard(selective bool) ReplyKeyboardHide {
	log.Println("NewHideKeyboard is deprecated, please use NewRemoveKeyboard")

	return ReplyKeyboardHide{
		HideKeyboard: true,
		Selective:    selective,
	}
}

// NewRemoveKeyboard hides the keyboard, with the option for being selective
// or hiding for everyone.
func NewRemoveKeyboard(selective bool) ReplyKeyboardRemove {
	return ReplyKeyboardRemove{
		RemoveKeyboard: true,
		Selective:      selective,
	}
}

// NewKeyboardButton creates a regular keyboard button.
func NewKeyboardButton(text string) KeyboardButton {
	return KeyboardButton{
		Text: text,
	}
}

// NewKeyboardButtonContact creates a keyboard button that requests
// user contact information upon click.
func NewKeyboardButtonContact(text string) KeyboardButton {
	return KeyboardButton{
		Text:           text,
		RequestContact: true,
	}
}

// NewKeyboardButtonLocation creates a keyboard button that requests
// user location information upon click.
func NewKeyboardButtonLocation(text string) KeyboardButton {
	return KeyboardButton{
		Text:            text,
		RequestLocation: true,
	}
}

// NewKeyboardButtonRow creates a row of keyboard buttons.
func NewKeyboardButtonRow(buttons ...*KeyboardButton) []*KeyboardButton {
	var row []*KeyboardButton

	row = append(row, buttons...)

	return row
}

// NewReplyKeyboard creates a new regular keyboard with sane defaults.
func NewReplyKeyboard(rows ...[]KeyboardButton) ReplyKeyboardMarkup {
	var keyboard [][]KeyboardButton

	keyboard = append(keyboard, rows...)

	return ReplyKeyboardMarkup{
		ResizeKeyboard: true,
		Keyboard:       keyboard,
	}
}

// NewInlineKeyboardButtonData creates an inline keyboard button with text
// and data for a callback.
func NewInlineKeyboardButtonData(text, data string) InlineKeyboardButton {
	return InlineKeyboardButton{
		Text:         text,
		CallbackData: &data,
	}
}

// NewInlineKeyboardButtonURL creates an inline keyboard button with text
// which goes to a URL.
func NewInlineKeyboardButtonURL(text, url string) InlineKeyboardButton {
	return InlineKeyboardButton{
		Text: text,
		URL:  &url,
	}
}

// NewInlineKeyboardButton creates an inline keyboard button with text
// user will send message after
func NewInlineKeyboardButton(text, url string) InlineKeyboardButton {
	return InlineKeyboardButton{
		Text: text,
	}
}

// NewInlineKeyboardButtonSwitch creates an inline keyboard button with
// text which allows the user to switch to a chat or return to a chat.
func NewInlineKeyboardButtonSwitch(text, sw string) InlineKeyboardButton {
	return InlineKeyboardButton{
		Text:              text,
		SwitchInlineQuery: &sw,
	}
}

// NewInlineKeyboardButtonPay creates an inline keyboard button with
// pay functionality.
func NewInlineKeyboardButtonPay(text string) InlineKeyboardButton {
	return InlineKeyboardButton{
		Text: text,
		Pay:  true,
	}
}

// NewInlineKeyboardRow creates an inline keyboard row with buttons.
func NewInlineKeyboardRow(buttons ...*InlineKeyboardButton) []*InlineKeyboardButton {
	var row []*InlineKeyboardButton

	row = append(row, buttons...)

	return row
}

// NewInlineKeyboardMarkup creates a new inline keyboard.
func NewInlineKeyboardMarkup(rows ...[]InlineKeyboardButton) InlineKeyboardMarkup {
	var keyboard [][]InlineKeyboardButton

	keyboard = append(keyboard, rows...)

	return InlineKeyboardMarkup{
		InlineKeyboard: keyboard,
	}
}

// NewCallback creates a new callback message.
func NewCallback(id, text string) CallbackConfig {
	return CallbackConfig{
		CallbackQueryID: id,
		Text:            text,
		ShowAlert:       false,
	}
}

// NewCallbackWithAlert creates a new callback message that alerts
// the user.
func NewCallbackWithAlert(id, text string) CallbackConfig {
	return CallbackConfig{
		CallbackQueryID: id,
		Text:            text,
		ShowAlert:       true,
	}
}

// NewInvoice creates a new Invoice request to the user.
func NewInvoice(chatID int64, title, description, payload, providerToken, startParameter, currency string, prices *[]LabeledPrice) InvoiceConfig {
	return InvoiceConfig{
		BaseChat:       BaseChat{ChatID: chatID},
		Title:          title,
		Description:    description,
		Payload:        payload,
		ProviderToken:  providerToken,
		StartParameter: startParameter,
		Currency:       currency,
		Prices:         prices}
}

// NewSetChatPhotoUpload creates a new chat photo uploader.
//
// chatID is where to send it, file is a string path to the file,
// FileReader, or FileBytes.
//
// Note that you must send animated GIFs as a document.
func NewSetChatPhotoUpload(chatID int64, file interface{}) SetChatPhotoConfig {
	return SetChatPhotoConfig{
		BaseFile: BaseFile{
			BaseChat:    BaseChat{ChatID: chatID},
			File:        file,
			UseExisting: false,
		},
	}
}

// NewSetChatPhotoShare shares an existing photo.
// You may use this to reshare an existing photo without reuploading it.
//
// chatID is where to send it, fileID is the ID of the file
// already uploaded.
func NewSetChatPhotoShare(chatID int64, fileID string) SetChatPhotoConfig {
	return SetChatPhotoConfig{
		BaseFile: BaseFile{
			BaseChat:    BaseChat{ChatID: chatID},
			FileID:      fileID,
			UseExisting: true,
		},
	}
}

// NewProduct creates a new Product.
// Requires both productID and eshopToken.
func NewProduct(chatID int64, productID int64, eshopToken string) ProductConfig {
	return ProductConfig{
		BaseChat: BaseChat{
			ChatID:           chatID,
			ReplyToMessageID: 0,
		},
		ProductID:  productID,
		EshopToken: eshopToken,
	}
}

// NewProductWithPeer creates a new Product.
// Requires both productID and eshopToken.
//
// This method only available for bots with privileges.
func NewProductWithPeer(peerID int32, productID int64, eshopToken string) ProductConfig {
	return ProductConfig{
		BaseChat: BaseChat{
			PeerID:           peerID,
			ReplyToMessageID: 0,
		},
		ProductID:  productID,
		EshopToken: eshopToken,
	}
}

// NewProductCart creates a new ProductCart.
// Requires a token.
func NewProductCart(chatID int64, cartToken string) ProductCartConfig {
	return ProductCartConfig{
		BaseChat: BaseChat{
			ChatID:           chatID,
			ReplyToMessageID: 0,
		},
		CartToken: cartToken,
	}
}

// NewProductCartWithPeer creates a new ProductCart.
// Requires a token.
//
// This method only available for bots with privileges.
func NewProductCartWithPeer(peerID int32, cartToken string) ProductCartConfig {
	return ProductCartConfig{
		BaseChat: BaseChat{
			PeerID:           peerID,
			ReplyToMessageID: 0,
		},
		CartToken: cartToken,
	}
}

// NewProductOrder creates a new ProductOrder.
// Requires a token.
func NewProductOrder(chatID int64, token string) ProductOrderConfig {
	return ProductOrderConfig{
		BaseChat: BaseChat{
			ChatID:           chatID,
			ReplyToMessageID: 0,
			ReplyMarkup: ReplyKeyboardMarkup{
				Keyboard:        [][]KeyboardButton{},
				ResizeKeyboard:  false,
				OneTimeKeyboard: false,
				Selective:       false,
			},
		},
		CartToken: token,
	}
}

// NewProductOrderWithPeer creates a new ProductOrder.
// Requires a token, peerID is the id of the user to send it to.
//
// This method only works for bots with privileges.
func NewProductOrderWithPeer(peerID int32, token string) ProductOrderConfig {
	return ProductOrderConfig{
		BaseChat: BaseChat{
			PeerID:           peerID,
			ReplyToMessageID: 0,
		},
		CartToken: token,
	}
}

// NewBotCommand creates a new BotCommand
func NewBotCommand(command, description string) BotCommand {
	return BotCommand{
		Command:     command,
		Description: description,
	}
}

// NewPhotoWithUrl creates a new photo by url.
//
// chatID is where to send it
//
// Note that you must send animated GIFs as a document.
func NewPhotoUrl(chatID int64, src string) PhotoUrlConfig {
	return PhotoUrlConfig{
		BaseChat: BaseChat{
			ChatID: chatID,
		},
		Url: src,
	}
}

// NewPhotoUrlWithPeer creates a new photo by url.
//
// peerID is the id of the user to send it to.
//
// Note that you must send animated GIFs as a document.
// This method only work for bots with privileges.
func NewPhotoUrlWithPeer(peerID int32, src string) PhotoUrlConfig {
	return PhotoUrlConfig{
		BaseChat: BaseChat{
			PeerID: peerID,
		},
		Url: src,
	}
}
