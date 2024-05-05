package presenter

import "github.com/saitamau-maximum/meline/domain/entity"

type NotifyMeta struct {
	TypeID uint64 `json:"type_id"`
}

type NotifyMessage struct {
	ID        string `json:"id"`
	User      *User  `json:"user"`
	Content   string `json:"content"`
	ChannelID uint64 `json:"channel_id"`
}

type NotifyMessageResponse struct {
	NotifyMeta NotifyMeta     `json:"notify_meta"`
	Message    *NotifyMessage `json:"message"`
}

type INotifyPresenter interface {
	GenerateNotifyMessageResponse(message *entity.Message) *NotifyMessageResponse
}
