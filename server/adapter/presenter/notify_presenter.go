package presenter

import (
	"github.com/saitamau-maximum/meline/config"
	"github.com/saitamau-maximum/meline/domain/entity"
	"github.com/saitamau-maximum/meline/usecase/presenter"
)

type NotifyPresenter struct{}

func NewNotifyPresenter() presenter.INotifyPresenter {
	return &NotifyPresenter{}
}

func (p *NotifyPresenter) GenerateNotifyMessageResponse(message *entity.Message) *presenter.NotifyMessageResponse {
	return &presenter.NotifyMessageResponse{
		NotifyMeta: presenter.NotifyMeta{
			TypeID: config.NOTIFY_MESSAGE,
		},
		Payload: presenter.Payload{
			Message: &presenter.Message{
				ID: message.ID,
				User: &presenter.User{
					ID:       message.User.ID,
					Name:     message.User.Name,
					ImageURL: message.User.ImageURL,
				},
				Content:        message.Content,
				ReplyToMessage: nil,
				CreatedAt:      message.CreatedAt.String(),
				UpdatedAt:      message.UpdatedAt.String(),
			},
			ChannelID: message.ChannelID,
		},
	}
}
