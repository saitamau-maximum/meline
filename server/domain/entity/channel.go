package entity

import (
	"time"
)

type Channel struct {
	ID        uint64
	Name      string
	ChannelID uint64
	Channels  []*Channel
	Users     []*User
	Messages  []*Message
	CreatedAt time.Time
	DeletedAt time.Time
}

func NewChannelEntity(id uint64, name string, channelID uint64, channels []*Channel, users []*User, messages []*Message, createdAt, deletedAt time.Time) *Channel {
	return &Channel{
		ID:        id,
		Name:      name,
		ChannelID: channelID,
		Channels:  channels,
		Users:     users,
		Messages:  messages,
		CreatedAt: createdAt,
		DeletedAt: deletedAt,
	}
}
