package presenter

import "github.com/saitamau-maximum/meline/domain/entity"

type Message struct {
	ID             string   `json:"id"`
	User           *User    `json:"user"`
	Content        string   `json:"content"`
	ReplyToMessage *ReplyToMessage `json:"reply_to_message"`
	CreatedAt      string   `json:"created_at"`
	UpdatedAt      string   `json:"updated_at"`
}

type Thread struct {
	ID             string     `json:"id"`
	User           *User      `json:"user"`
	Content        string     `json:"content"`
	Comments       []*Comments `json:"comments"`
	CreatedAt      string     `json:"created_at"`
	UpdatedAt      string     `json:"updated_at"`
}

type ReplyToMessage struct {
	ID        string `json:"id"`
	User      *User  `json:"user"`
	Content   string `json:"content"`
}

type Comments struct {
	ID      string `json:"id"`
	User    *User  `json:"user"`
	Content string `json:"content"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type GetMessageByIDResponse struct {
	Thread *Thread `json:"thread"`
}

type IMessagePresenter interface {
	GenerateGetMessageByIDResponse(message *entity.Message) *GetMessageByIDResponse
}
