package model

import (
	"time"

	"github.com/saitamau-maximum/meline/domain/entity"
)

type Channel struct {
	ID        uint64    `bun:"id,pk,autoincrement"`
	Name      string    `bun:"name,notnull"`
	ParentChannelID uint64    `bun:"parent_channel_id"`
	Channels  []*Channel `bun:"rel:has-many,join:id=parent_channel_id"`
	Users     []*User   `bun:"m2m:channel_users,join:Channel=User"`
	Messages  []*Message `bun:"rel:has-many,join:id=channel_id"`
	CreatedAt time.Time `bun:"created_at,notnull,default:current_timestamp"`
	DeletedAt time.Time `bun:"deleted_at,default:null"`
}

func (c *Channel) ToChannelEntity() *entity.Channel {
	entitiedUsers := make([]*entity.User, len(c.Users))
	for i, u := range c.Users {
		entitiedUsers[i] = u.ToUserEntity()
	}

	entitiedChannels := make([]*entity.Channel, len(c.Channels))
	for i, ch := range c.Channels {
		entitiedChannels[i] = ch.ToChannelEntity()
	}

	entitiedMessages := make([]*entity.Message, len(c.Messages))
	for i, m := range c.Messages {
		entitiedMessages[i] = m.ToMessageEntity()
	}

	return entity.NewChannelEntity(c.ID, c.Name, c.ParentChannelID, entitiedChannels, entitiedUsers, entitiedMessages, c.CreatedAt, c.DeletedAt)
}
