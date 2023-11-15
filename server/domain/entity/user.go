package entity

import (
	"time"
)

type User struct {
	ID uint64 `json:"id"`
	ProviderID string `json:"provider_id"`
	Name string `json:"name"`
	ImageURL string `json:"image_url"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}

func NewUserEntity(id uint64, providerID, name string, createdAt, updatedAt, deletedAt time.Time) *User {
	return &User{
		ID: id,
		ProviderID: providerID,
		Name: name,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
		DeletedAt: deletedAt,
	}
}
