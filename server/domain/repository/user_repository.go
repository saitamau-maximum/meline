package repository

import (
	"context"

	"github.com/saitamau-maximum/meline/models"
)

type IUserRepository interface {
	FindByID(ctx context.Context, id uint64) (*model.User, error)
	FindByGithubID(ctx context.Context, githubID string) (*model.User, error)
	Insert(ctx context.Context, githubID, name string) error
	Update(ctx context.Context, id uint64, githubID, name string) error
	Delete(ctx context.Context, id uint64) error
}
