package entity

import (
	"time"
)

type Message struct {
	ID             string
	ChannelID      uint64
	Channel        *Channel
	UserID         uint64
	User           *User
	ReplyToMessage []*Message
	Content        string
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      time.Time
}

func NewMessageEntity(id string, channelID uint64, channel *Channel, userID uint64, user *User, replyToMessage []*Message, content string, createdAt, updatedAt, deletedAt time.Time) *Message {
	return &Message{
		ID:             id,
		ChannelID:      channelID,
		Channel:        channel,
		UserID:         userID,
		User:           user,
		ReplyToMessage: replyToMessage,
		Content:        content,
		CreatedAt:      createdAt,
		UpdatedAt:      updatedAt,
		DeletedAt:      deletedAt,
	}
}
