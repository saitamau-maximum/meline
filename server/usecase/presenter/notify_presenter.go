package presenter

import "github.com/saitamau-maximum/meline/domain/entity"

type NotifyMeta struct {
	TypeID uint64 `json:"type_id"`
}

type Payload struct {
	Message   *Message `json:"message"`
	ChannelID uint64   `json:"channel_id"`
}

type NotifyMessageResponse struct {
	NotifyMeta NotifyMeta `json:"notify_meta"`
	Payload    Payload    `json:"payload"`
}

type INotifyPresenter interface {
	GenerateNotifyMessageResponse(message *entity.Message) *NotifyMessageResponse
}
