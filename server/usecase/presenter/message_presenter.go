package presenter

import "github.com/saitamau-maximum/meline/domain/entity"

type Message struct {
	ID             string          `json:"id"`
	User           *User           `json:"user"`
	Content        string          `json:"content"`
	ReplyToMessage *ReplyToMessage `json:"reply_to_message"`
	CreatedAt      string          `json:"created_at"`
	UpdatedAt      string          `json:"updated_at"`
}

type ReplyToMessage struct {
	ID      string `json:"id"`
	User    *User  `json:"user"`
	Content string `json:"content"`
}

type GetMessagesByChannelIDResponse struct {
	Messages []*Message `json:"messages"`
}

type CreateMessageResponse struct {
	Message   *Message `json:"message"`
	ChannelID uint64   `json:"channel_id"`
}

type IMessagePresenter interface {
	GenerateGetMessagesByChannelIDResponse(messages []*entity.Message) *GetMessagesByChannelIDResponse
	GenerateCreateMessageResponse(message *entity.Message) *CreateMessageResponse
}
