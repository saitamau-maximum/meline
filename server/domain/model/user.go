package model

import "time"

type User struct {
	ID uint64 `json:"id", bun:"id,primary_key,auto_increment"`
	GitHubID string `json:"github_id", bun:"github_id,unique"`
	Name string `json:"name", bun:"name"`
	CreatedAt time.Time `json:"created_at", bun:"created_at, default:current_timestamp"`
	UpdatedAt time.Time `json:"updated_at", bun:"updated_at, default:current_timestamp"`
	DeletedAt time.Time `json:"deleted_at", bun:"deleted_at, nullable, default:null"`
}