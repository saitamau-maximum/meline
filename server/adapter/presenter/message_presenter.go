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
		if len(message.ReplyToMessage) > 0 {
			replyToMessage.ID = message.ReplyToMessage[0].ID
			replyToMessage.User = &presenter.User{
				ID:       message.ReplyToMessage[0].User.ID,
				Name:     message.ReplyToMessage[0].User.Name,
				ImageURL: message.ReplyToMessage[0].User.ImageURL,
			}
			replyToMessage.Content = message.ReplyToMessage[0].Content
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
