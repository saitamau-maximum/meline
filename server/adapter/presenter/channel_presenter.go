package presenter

import (
	"strconv"

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
			ID:   strconv.FormatUint(channel.ID, 10),
			Name: channel.Name,
		})
	}

	return channelsResponse
}

func (p *ChannelPresenter) GenerateGetChannelByIdResponse(channel *entity.Channel) *presenter.GetChannelByIdResponse {
	childChannels := make([]*presenter.Channel, 0)
	for _, childChannel := range channel.ChildChannels {
		childChannels = append(childChannels, &presenter.Channel{
			ID:   strconv.FormatUint(childChannel.ID, 10),
			Name: childChannel.Name,
		})
	}

	users := make([]*presenter.User, 0)
	for _, user := range channel.Users {
		users = append(users, &presenter.User{
			ID:       strconv.FormatUint(user.ID, 10),
			Name:     user.Name,
			ImageURL: user.ImageURL,
		})
	}

	return &presenter.GetChannelByIdResponse{
		Channel: &presenter.ChannelDetail{
			ID:       strconv.FormatUint(channel.ID, 10),
			Name:     channel.Name,
			Users:    users,
			Channels: childChannels,
		},
	}
}

func (p *ChannelPresenter) GenerateGetChannelsByNameResponse(channels []*entity.Channel) *presenter.GetChannelsByNameResponse {
	channelsResponse := &presenter.GetChannelsByNameResponse{
		Channels: []*presenter.Channel{},
	}
	for _, channel := range channels {
		channelsResponse.Channels = append(channelsResponse.Channels, &presenter.Channel{
			ID:   strconv.FormatUint(channel.ID, 10),
			Name: channel.Name,
		})
	}

	return channelsResponse
}
