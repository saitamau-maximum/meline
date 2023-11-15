package repository

import (
	"context"

	"github.com/saitamau-maximum/meline/models"
)

type IUserRepository interface {
	FindByID(ctx context.Context, id uint64) (*model.User, error)
	FindByGithubID(ctx context.Context, githubID string) (*model.User, error)
	Create(ctx context.Context, user *model.User) error
}
