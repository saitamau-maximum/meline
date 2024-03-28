package model

import (
	"time"

	"github.com/saitamau-maximum/meline/domain/entity"
)

type Channel struct {
	ID        uint64     `bun:"id,pk,autoincrement"`
	Name      string     `bun:"name,notnull"`
	Users     []*User    `bun:"m2m:channel_users,join:Channel=User"`
	Messages  []*Message `bun:"rel:has-many,join:id=channel_id"`
	CreatedAt time.Time  `bun:"created_at,notnull,default:current_timestamp"`
	DeletedAt time.Time  `bun:"deleted_at,default:null"`
}

func (c *Channel) ToChannelEntity() *entity.Channel {
	entitiedUsers := make([]*entity.User, len(c.Users))
	for i, u := range c.Users {
		entitiedUsers[i] = u.ToUserEntity()
	}

	return entity.NewChannelEntity(c.ID, c.Name, entitiedUsers, c.CreatedAt, c.DeletedAt)
}
