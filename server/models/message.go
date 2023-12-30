package model

import (
	"time"

	"github.com/saitamau-maximum/meline/domain/entity"
	"github.com/saitamau-maximum/meline/utils"
)

type Message struct {
	ID        string     `bun:"id,pk"`
	ChannelID uint64     `bun:"channel_id,notnull"`
	UserID    uint64     `bun:"user_id,notnull"`
	Replys    []*Message `bun:"rel:has-many,join:reply_to_id=id"`
	ReplyToID string     `bun:"reply_to_id,default:null"`
	Content   string     `bun:"content,notnull,type:varchar(2000)"`
	CreatedAt time.Time  `bun:"created_at,notnull,default:current_timestamp"`
	UpdatedAt time.Time  `bun:"updated_at,notnull,default:current_timestamp"`
	DeletedAt time.Time  `bun:"deleted_at,default:null"`
}

func (m *Message) ToMessageEntity() *entity.Message {
	replys := make([]*entity.Message, len(m.Replys))
	for i, r := range m.Replys {
		replys[i] = r.ToMessageEntity()
	}

	return entity.NewMessageEntity(m.ID, m.ChannelID, m.UserID, replys, m.ReplyToID, m.Content, m.CreatedAt, m.UpdatedAt, m.DeletedAt)
}

func NewMessageModel(channelID uint64, userID uint64, replys []*Message, replyToID string, content string, createdAt time.Time, updatedAt time.Time, deletedAt time.Time) *Message {
	return &Message{
		ID:        utils.GenerateUUID(),
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
