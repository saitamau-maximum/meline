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
	ReplyToMessage *Message
	ReplyToID      string
	Replys         []*Message
	Content        string
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      time.Time
}

func NewMessageEntity(id string, channelID uint64, channel *Channel, userID uint64, user *User, replyToMessage *Message, replyToID string, content string, createdAt time.Time, updatedAt time.Time, deletedAt time.Time) *Message {
	return &Message{
		ID:             id,
		ChannelID:      channelID,
		Channel:        channel,
		UserID:         userID,
		User:           user,
		ReplyToMessage: replyToMessage,
		ReplyToID:      replyToID,
		Content:        content,
		CreatedAt:      createdAt,
		UpdatedAt:      updatedAt,
		DeletedAt:      deletedAt,
	}
}
