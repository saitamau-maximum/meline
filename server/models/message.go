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
	ReplyToMessage *Message   `bun:"rel:belongs-to,join:reply_to_id=id"`
	Replys         []*Message `bun:"rel:has-many,join:id=reply_to_id"`
	ReplyToID      string     `bun:"reply_to_id,default:null"`
	Content        string     `bun:"content,notnull,type:varchar(2000)"`
	CreatedAt      time.Time  `bun:"created_at,notnull,default:current_timestamp"`
	UpdatedAt      time.Time  `bun:"updated_at,notnull,default:current_timestamp"`
	DeletedAt      time.Time  `bun:"deleted_at,default:null"`
}

func (m *Message) ToMessageEntity() *entity.Message {
	replys := make([]*entity.Message, len(m.Replys))
	for i, r := range m.Replys {
		replys[i] = r.ToMessageEntity()
	}

	return entity.NewMessageEntity(m.ID, m.ChannelID, m.Channel.ToChannelEntity(), m.UserID, m.User.ToUserEntity(), m.ReplyToMessage.ToMessageEntity(), m.ReplyToID, m.Content, m.CreatedAt, m.UpdatedAt, m.DeletedAt)
}

func NewMessageModel(channelID uint64, userID uint64, replyToID string, content string) *Message {
	return &Message{
		ID:        utils.GenerateUUID(),
		ChannelID: channelID,
		UserID:    userID,
		ReplyToID: replyToID,
		Content:   content,
	}
}
