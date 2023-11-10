package entity

import (
	"time"
)

type User struct {
	ID uint64 `json:"id"`
	GithubID string `json:"github_id"`
	Name string `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}

func NewUserEntity(id uint64, githubID, name string, createdAt, updatedAt, deletedAt time.Time) *User {
	return &User{
		ID: id,
		GithubID: githubID,
		Name: name,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
		DeletedAt: deletedAt,
	}
}
