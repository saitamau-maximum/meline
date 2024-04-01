package model

import (
	"time"

	"github.com/saitamau-maximum/meline/domain/entity"
)

type Channel struct {
	ID            uint64     `bun:"id,pk,autoincrement"`
	Name          string     `bun:"name,notnull"`
	Users         []*User    `bun:"m2m:channel_users,join:Channel=User"`
	Messages      []*Message `bun:"rel:has-many,join:id=channel_id"`
	ChildChannels []*Channel `bun:"m2m:channel_to_channels,join:ParentChannel=ChildChannel"`
	CreatedAt     time.Time  `bun:"created_at,notnull,default:current_timestamp"`
	DeletedAt     time.Time  `bun:"deleted_at,default:null"`
}

func (c *Channel) ToChannelEntity() *entity.Channel {
	entitiedUsers := make([]*entity.User, len(c.Users))
	for i, u := range c.Users {
		entitiedUsers[i] = u.ToUserEntity()
	}

	entitiedMessages := make([]*entity.Message, len(c.Messages))
	for i, m := range c.Messages {
		entitiedMessages[i] = m.ToMessageEntity()
	}

	entitiedChannels := make([]*entity.Channel, len(c.ChildChannels))
	for i, ch := range c.ChildChannels {
		entitiedChannels[i] = ch.ToChannelEntity()
	}

	return entity.NewChannelEntity(c.ID, c.Name, entitiedUsers, entitiedChannels, entitiedMessages, c.CreatedAt, c.DeletedAt)
}
