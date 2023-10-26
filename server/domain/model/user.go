package model

import "time"

type User struct {
	ID uint64 `bun:"id,pk,auto_increment,notnull"`
	GithubId string `bun:"github_id,unique,notnull"`
	Name string `bun:"name,notnull"`
	CreatedAt time.Time `bun:"created_at,notnull,default:current_timestamp"`
	UpdatedAt time.Time `bun:"updated_at,notnull,default:current_timestamp"`
	DeletedAt time.Time `bun:"deleted_at,nullable,default:null"`
}