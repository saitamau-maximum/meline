package model

import (
	"time"

	"github.com/saitamau-maximum/meline/domain/entity"
	"github.com/saitamau-maximum/meline/utils"
)

type Message struct {
	ID             string     `bun:"id,pk"`
	ChannelID      uint64     `bun:"channel_id,notnull"`
	Channel        *Channel   `bun:"rel:belongs-to,join:channel_id=id"`
	UserID         uint64     `bun:"user_id,notnull"`
	User           *User      `bun:"rel:belongs-to,join:user_id=id"`
	ReplyToMessage []*Message `bun:"m2m:message_to_messages,join:ChildMessage=ParentMessage"`
	Content        string     `bun:"content,notnull,type:varchar(2000)"`
	CreatedAt      time.Time  `bun:"created_at,notnull,default:current_timestamp"`
	UpdatedAt      time.Time  `bun:"updated_at,notnull,default:current_timestamp"`
	DeletedAt      time.Time  `bun:"deleted_at,soft_delete,nullzero"`
}

func (m *Message) ToMessageEntity() *entity.Message {
	var entitiedChannel *entity.Channel
	if m.Channel != nil {
		entitiedChannel = m.Channel.ToChannelEntity()
	}

	var entitiedUser *entity.User
	if m.User != nil {
		entitiedUser = m.User.ToUserEntity()
	}

	entitiedReplyToMessages := []*entity.Message{}
	if m.ReplyToMessage != nil {
		for _, r := range m.ReplyToMessage {
			entitiedReplyToMessages = append(entitiedReplyToMessages, r.ToMessageEntity())
		}
	}

	return entity.NewMessageEntity(m.ID, m.ChannelID, entitiedChannel, m.UserID, entitiedUser, entitiedReplyToMessages, m.Content, m.CreatedAt, m.UpdatedAt, m.DeletedAt)
}

func NewMessageModel(channelID uint64, userID uint64, content string) *Message {
	return &Message{
		ID:        utils.GenerateUUID(),
		ChannelID: channelID,
		UserID:    userID,
		Content:   content,
	}
}
