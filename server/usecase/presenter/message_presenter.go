package presenter

import "github.com/saitamau-maximum/meline/domain/entity"

type Message struct {
	ID             string   `json:"id"`
	User           *User    `json:"user"`
	Content        string   `json:"content"`
	ReplyToMessage *Message `json:"reply_to_message"`
}

type MessageDetail struct {
	ID             string     `json:"id"`
	User           *User      `json:"user"`
	Channel        *Channel   `json:"channel"`
	Content        string     `json:"content"`
	ReplyToMessage *Message   `json:"reply_to_message"`
	Replys         []*Message `json:"replys"`
}

type GetMessageByIDResponse struct {
	Message *MessageDetail `json:"message"`
}

type GetMessagesByChannelIDResponse struct {
	Messages []*Message `json:"messages"`
}

type IMessagePresenter interface {
	GenerateGetMessageByIDResponse(message *entity.Message) *GetMessageByIDResponse
	GenerateGetMessagesByChannelIDResponse(messages []*entity.Message) *GetMessagesByChannelIDResponse
}
