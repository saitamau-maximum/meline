package model

import "time"

type Notify struct {
	ID        uint64    `bun:"id,pk,autoincrement"`
	UserID    uint64    `bun:"user_id,notnull"`
	MessageID string    `bun:"message_id"`
	TypeID    uint64    `bun:"type_id,notnull"`
	Message   *Message  `bun:"rel:has-one,join:message_id=id"`
	CreatedAt time.Time `bun:"created_at,notnull,default:current_timestamp"`
	UpdatedAt time.Time `bun:"updated_at,notnull,default:current_timestamp"`
	DeletedAt time.Time `bun:"deleted_at,soft_delete,nullzero"`
}
