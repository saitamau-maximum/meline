package entity

import (
	"time"
)

type Message struct {
	ID        string
	ChannelID uint64
	UserID    uint64
	Replys    []*Message
	ReplyToID string
	Content   string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

func NewMessageEntity(id string, channelID uint64, userID uint64, replys []*Message, replyToID string, content string, createdAt time.Time, updatedAt time.Time, deletedAt time.Time) *Message {
	return &Message{
		ID:        id,
		ChannelID: channelID,
		UserID:    userID,
		Replys:    replys,
		ReplyToID: replyToID,
		Content:   content,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
		DeletedAt: deletedAt,
	}
}
