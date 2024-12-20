package presenter

import (
	"github.com/saitamau-maximum/meline/domain/entity"
)

type Channel struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type ChannelDetail struct {
	ID       string     `json:"id"`
	Name     string     `json:"name"`
	Users    []*User    `json:"users"`
	Channels []*Channel `json:"channels"`
}

type GetAllChannelsResponse struct {
	Channels []*Channel `json:"channels"`
}

type GetChannelByIdResponse struct {
	Channel *ChannelDetail `json:"channel"`
}

type GetChannelsByNameResponse struct {
	Channels []*Channel `json:"channels"`
}

type CreateChannelResponse struct {
	ID string `json:"id"`
}

type UpdateChannelResponse struct {
	ID string `json:"id"`
}

type IChannelPresenter interface {
	GenerateGetAllChannelsResponse(channels []*entity.Channel) *GetAllChannelsResponse
	GenerateGetChannelByIdResponse(channel *entity.Channel) *GetChannelByIdResponse
	GenerateGetChannelsByNameResponse(channels []*entity.Channel) *GetChannelsByNameResponse
}
