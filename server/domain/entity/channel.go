package entity

import (
	"time"
)

type Channel struct {
	ID        uint64
	Name      string
	Users     []*User
	CreatedAt time.Time
	DeletedAt time.Time
}

func NewChannelEntity(id uint64, name string, Users []*User, createdAt, deletedAt time.Time) *Channel {
	return &Channel{
		ID:        id,
		Name:      name,
		Users:     Users,
		CreatedAt: createdAt,
		DeletedAt: deletedAt,
	}
}
