package entity

import (
	"time"
)

type User struct {
	ID uint64 
	GithubID string 
	Name string 
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
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