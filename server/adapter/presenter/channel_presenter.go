package presenter

import (
	"github.com/saitamau-maximum/meline/domain/entity"
	"github.com/saitamau-maximum/meline/usecase/presenter"
)

type ChannelPresenter struct{}

func NewChannelPresenter() presenter.IChannelPresenter {
	return &ChannelPresenter{}
}

func (p *ChannelPresenter) GenerateGetAllChannelsResponse(channels []*entity.Channel) *presenter.GetAllChannelsResponse {
	channelsResponse := &presenter.GetAllChannelsResponse{
		Channels: []*presenter.Channel{},
	}
	for _, channel := range channels {
		channelsResponse.Channels = append(channelsResponse.Channels, &presenter.Channel{
			ID:   channel.ID,
			Name: channel.Name,
		})
	}

	return channelsResponse
}

func (p *ChannelPresenter) GenerateGetChannelByIdResponse(channel *entity.Channel) *presenter.GetChannelByIdResponse {
	users := make([]*presenter.User, len(channel.Users))
	for _, user := range channel.Users {
		users = append(users, &presenter.User{
			ID:       user.ID,
			Name:     user.Name,
			ImageURL: user.ImageURL,
		})
	}

	messages := make([]*presenter.Message, len(channel.Messages))
	for _, message := range channel.Messages {
		replyToMessage := &presenter.ReplyToMessage{}
		if message.ReplyToMessage != nil {
			replyToMessage = &presenter.ReplyToMessage{
				ID:        message.ReplyToMessage.ID,
				User:      &presenter.User{
					ID:       message.ReplyToMessage.User.ID,
					Name:     message.ReplyToMessage.User.Name,
					ImageURL: message.ReplyToMessage.User.ImageURL,
				},
				Content:   message.ReplyToMessage.Content,
			}
		}

		messages = append(messages, &presenter.Message{
			ID:        message.ID,
			User:      &presenter.User{
				ID:       message.User.ID,
				Name:     message.User.Name,
				ImageURL: message.User.ImageURL,
			},
			ReplyToMessage: replyToMessage,
			Content:   message.Content,
			CreatedAt: message.CreatedAt.String(),
			UpdatedAt: message.UpdatedAt.String(),
		})
	}

	return &presenter.GetChannelByIdResponse{
		Channel: &presenter.ChannelDetail{
			Name:  channel.Name,
			Users: users,
			Messages: messages,
		},
	}
}

func (p *ChannelPresenter) GenerateGetChannelsByNameResponse(channels []*entity.Channel) *presenter.GetChannelsByNameResponse {
	channelsResponse := &presenter.GetChannelsByNameResponse{
		Channels: []*presenter.Channel{},
	}
	for _, channel := range channels {
		channelsResponse.Channels = append(channelsResponse.Channels, &presenter.Channel{
			ID:   channel.ID,
			Name: channel.Name,
		})
	}

	return channelsResponse
}
