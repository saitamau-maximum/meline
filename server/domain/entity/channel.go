package entity

import (
	"time"
)

type Channel struct {
	ID            uint64
	Name          string
	Messages      []*Message
	ChildChannels []*Channel
	Users         []*User
	CreatedAt     time.Time
	DeletedAt     time.Time
}

func NewChannelEntity(id uint64, name string, users []*User, childChannels []*Channel, messages []*Message, createdAt time.Time, deletedAt time.Time) *Channel {
	return &Channel{
		ID:            id,
		Name:          name,
		Messages:      messages,
		ChildChannels: childChannels,
		Users:         users,
		CreatedAt:     createdAt,
		DeletedAt:     deletedAt,
	}
}
