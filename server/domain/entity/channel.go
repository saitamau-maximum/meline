package entity

import (
	"time"
)

type Channel struct {
	ID        uint64
	Name      string
	Users     []*User
	Messages  []*Message
	CreatedAt time.Time
	DeletedAt time.Time
}

func NewChannelEntity(id uint64, name string, users []*User, messages []*Message, createdAt, deletedAt time.Time) *Channel {
	return &Channel{
		ID:        id,
		Name:      name,
		Users:     users,
		Messages:  messages,
		CreatedAt: createdAt,
		DeletedAt: deletedAt,
	}
}
