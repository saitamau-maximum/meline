package model

import (
	"time"

	"github.com/saitamau-maximum/meline/domain/entity"
)

type User struct {
	ID         uint64     `bun:"id,pk,autoincrement"`
	ProviderID string     `bun:"provider_id,unique"`
	Name       string     `bun:"name,notnull"`
	ImageURL   string     `bun:"image_url"`
	Channels   []*Channel `bun:"m2m:channel_users,join:User=Channel"`
	CreatedAt  time.Time  `bun:"created_at,notnull,default:current_timestamp"`
	UpdatedAt  time.Time  `bun:"updated_at,notnull,default:current_timestamp"`
	DeletedAt  time.Time  `bun:"deleted_at,soft_delete"`
}

func (u *User) ToUserEntity() *entity.User {
	return entity.NewUserEntity(u.ID, u.ProviderID, u.Name, u.ImageURL, u.CreatedAt, u.UpdatedAt, u.DeletedAt)
}
