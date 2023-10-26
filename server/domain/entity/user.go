package entity

import "time"

type User struct {
	ID uint64 
	GithubID string 
	Name string 
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}