package presenter

import (
	"github.com/saitamau-maximum/meline/domain/entity"
	"github.com/saitamau-maximum/meline/usecase/presenter"
)

type MessagePresenter struct{}

func NewMessagePresenter() presenter.IMessagePresenter {
	return &MessagePresenter{}
}

func (p *MessagePresenter) GenerateGetMessagesByChannelIDResponse(messages []*entity.Message) *presenter.GetMessagesByChannelIDResponse {
	messagesResponse := &presenter.GetMessagesByChannelIDResponse{
		Messages: []*presenter.Message{},
	}
	for _, message := range messages {
		replyToMessage := &presenter.ReplyToMessage{}
		if message.ReplyToMessage != nil {
			replyToMessage.ID = message.ReplyToMessage.ID
			replyToMessage.User = &presenter.User{
				ID:       message.ReplyToMessage.User.ID,
				Name:     message.ReplyToMessage.User.Name,
				ImageURL: message.ReplyToMessage.User.ImageURL,
			}
			replyToMessage.Content = message.ReplyToMessage.Content
		}
		messagesResponse.Messages = append(messagesResponse.Messages, &presenter.Message{
			ID: message.ID,
			User: &presenter.User{
				ID:       message.User.ID,
				Name:     message.User.Name,
				ImageURL: message.User.ImageURL,
			},
			Content:        message.Content,
			ReplyToMessage: replyToMessage,
			CreatedAt:      message.CreatedAt.String(),
			UpdatedAt:      message.UpdatedAt.String(),
		})
	}

	return messagesResponse
}
