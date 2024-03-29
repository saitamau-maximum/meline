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
		replyToMessages := make([]*presenter.ReplyToMessage, 0)
		for _, replyToMessage := range message.ReplyToMessage {
			replyToMessages = append(replyToMessages, &presenter.ReplyToMessage{
				ID: replyToMessage.ID,
				User: &presenter.User{
					ID:       replyToMessage.User.ID,
					Name:     replyToMessage.User.Name,
					ImageURL: replyToMessage.User.ImageURL,
				},
				Content: replyToMessage.Content,
			})
		}
		messagesResponse.Messages = append(messagesResponse.Messages, &presenter.Message{
			ID: message.ID,
			User: &presenter.User{
				ID:       message.User.ID,
				Name:     message.User.Name,
				ImageURL: message.User.ImageURL,
			},
			Content:        message.Content,
			ReplyToMessage: replyToMessages,
			CreatedAt:      message.CreatedAt.String(),
			UpdatedAt:      message.UpdatedAt.String(),
		})
	}

	return messagesResponse
}
