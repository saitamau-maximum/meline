package presenter

import (
	"github.com/saitamau-maximum/meline/domain/entity"
)

type Channel struct {
	ID 		  uint64  `json:"id"`
	Name 	  string  `json:"name"`
}

type ChannelDetail struct {
	Name 	  string  `json:"name"`
	Users 	  []*User `json:"users"`
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
	ID uint64 `json:"id"`
}

type UpdateChannelResponse struct {
	ID uint64 `json:"id"`
}

type IChannelPresenter interface {
	GenerateGetAllChannelsResponse(channels []*entity.Channel) *GetAllChannelsResponse
	GenerateGetChannelByIdResponse(channel *entity.Channel) *GetChannelByIdResponse
	GenerateGetChannelsByNameResponse(channels []*entity.Channel) *GetChannelsByNameResponse
}
