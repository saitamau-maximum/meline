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
	channelsResponse := &presenter.GetAllChannelsResponse{}
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

	return &presenter.GetChannelByIdResponse{
		Channel: &presenter.ChannelDetail{
			Name:  channel.Name,
			Users: users,
		},
	}
}

func (p *ChannelPresenter) GenerateGetChannelsByNameResponse(channels []*entity.Channel) *presenter.GetChannelsByNameResponse {
	channelsResponse := &presenter.GetChannelsByNameResponse{}
	for _, channel := range channels {
		channelsResponse.Channels = append(channelsResponse.Channels, &presenter.Channel{
			ID:   channel.ID,
			Name: channel.Name,
		})
	}

	return channelsResponse
}
