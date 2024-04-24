package entity

import (
	"time"
)

type Message struct {
	ID               string
	ChannelID        uint64
	Channel          Channel
	UserID           uint64
	User             User
	ReplyToMessageID string
	ReplyToMessage   *Message
	Content          string
	CreatedAt        time.Time
	UpdatedAt        time.Time
	DeletedAt        *time.Time
}

func NewMessageEntity(id string, channelID uint64, channel Channel, userID uint64, user User, replyToMessageID string, replyToMessage *Message, content string, createdAt, updatedAt time.Time, deletedAt *time.Time) *Message {
	return &Message{
		ID:               id,
		ChannelID:        channelID,
		Channel:          channel,
		UserID:           userID,
		User:             user,
		ReplyToMessageID: replyToMessageID,
		ReplyToMessage:   replyToMessage,
		Content:          content,
		CreatedAt:        createdAt,
		UpdatedAt:        updatedAt,
		DeletedAt:        deletedAt,
	}
}
