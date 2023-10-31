package repository

import (
	"context"

	"github.com/saitamau-maximum/meline/models"
)

type IUserRepository interface {
	FindByGithubID(ctx context.Context, githubID string) (*model.User, error)
	Insert(ctx context.Context, user *model.User) error
	Update(ctx context.Context, user *model.User) error
	Delete(ctx context.Context, id uint64) error
}
