package model

import (
	"time"

	"github.com/google/uuid"
	"github.com/saitamau-maximum/meline/domain/entity"
)

type Message struct {
	ID               string    `bun:"id,pk"`
	ChannelID        uint64    `bun:"channel_id,notnull"`
	Channel          *Channel  `bun:"rel:belongs-to,join:channel_id=id"`
	UserID           uint64    `bun:"user_id,notnull"`
	User             *User     `bun:"rel:belongs-to,join:user_id=id"`
	ReplyToMessageID string    `bun:"reply_to_message_id"`
	ReplyToMessage   *Message  `bun:"rel:belongs-to,join:reply_to_message_id=id"`
	Content          string    `bun:"content,notnull,type:varchar(2000)"`
	CreatedAt        time.Time `bun:"created_at,notnull,default:current_timestamp"`
	UpdatedAt        time.Time `bun:"updated_at,notnull,default:current_timestamp"`
	DeletedAt        time.Time `bun:"deleted_at,soft_delete,nullzero"`
}

func (m *Message) ToMessageEntity() *entity.Message {
	var entitiedChannel entity.Channel
	if m.Channel != nil {
		entitiedChannel = *(m.Channel.ToChannelEntity())
	}

	var entitiedUser entity.User
	if m.User != nil {
		entitiedUser = *(m.User.ToUserEntity())
	}

	var entitiedReplyToMessage *entity.Message = nil
	if m.ReplyToMessage != nil {
		entitiedReplyToMessage = m.ReplyToMessage.ToMessageEntity()
	}

	var deletedAt *time.Time = nil
	if !m.DeletedAt.IsZero() {
		deletedAt = &m.DeletedAt
	}

	return entity.NewMessageEntity(m.ID, m.ChannelID, entitiedChannel, m.UserID, entitiedUser, m.ReplyToMessageID, entitiedReplyToMessage, m.Content, m.CreatedAt, m.UpdatedAt, deletedAt)
}

func NewMessageModel(channelID uint64, userID uint64, content string) *Message {
	return &Message{
		ID:        uuid.New().String(),
		ChannelID: channelID,
		UserID:    userID,
		Content:   content,
	}
}
