package entity

import (
	"time"
)

type ChannelUsers struct {
	UserID    uint64
	User      *User
	ChannelID uint64
	Channel   *Channel
	JoinedAt  time.Time
}

func NewChannelUsersEntity(userID uint64, user *User, channelID uint64, channel *Channel, joinedAt time.Time) *ChannelUsers {
	return &ChannelUsers{
		UserID:    userID,
		User:      user,
		ChannelID: channelID,
		Channel:   channel,
		JoinedAt:  joinedAt,
	}
}
