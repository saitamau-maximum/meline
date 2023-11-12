package model

import (
	"time"

	"github.com/saitamau-maximum/meline/domain/entity"
)

type User struct {
	ID uint64 `bun:"id,pk,autoincrement"`
	GithubID string `bun:"github_id,unique,notnull"`
	Name string `bun:"name,notnull"`
	ImageURL string `bun:"image_url"`
	CreatedAt time.Time `bun:"created_at,notnull,default:current_timestamp"`
	UpdatedAt time.Time `bun:"updated_at,notnull,default:current_timestamp"`
	DeletedAt time.Time `bun:"deleted_at,default:null"`
}

func (u *User) ToUserEntity() *entity.User {
	return entity.NewUserEntity(u.ID, u.GithubID, u.Name, u.CreatedAt, u.UpdatedAt, u.DeletedAt)
}
