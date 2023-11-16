package repository

import (
	"context"

	model "github.com/saitamau-maximum/meline/models"
)

type IUserRepository interface {
	FindByID(ctx context.Context, id uint64) (*model.User, error)
	FindByProviderID(ctx context.Context, providerID string) (*model.User, error)
	Create(ctx context.Context, user *model.User) error
}
