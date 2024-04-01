package model

import (
	"time"
)

type ChannelUsers struct {
	UserID    uint64    `bun:"user_id,pk"`
	User      *User     `bun:"rel:belongs-to,join:user_id=id"`
	ChannelID uint64    `bun:"channel_id,pk"`
	Channel   *Channel  `bun:"rel:belongs-to,join:channel_id=id"`
	JoinedAt  time.Time `bun:"joined_at,notnull,default:current_timestamp"`
}
