package infra

import (
	"context"
	"sync"

	"github.com/uptrace/bun"

	"github.com/saitamau-maximum/meline/domain/repository"
	"github.com/saitamau-maximum/meline/models"
)

type UserRepository struct {
	db *bun.DB
	mu sync.RWMutex
}

func NewUserRepository(db *bun.DB) repository.IUserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) FindByGithubID(ctx context.Context, githubID string) (*model.User, error) {
	return nil, nil
}

func (r *UserRepository) Insert(ctx context.Context, user *model.User) error {
	return nil
}

func (r *UserRepository) Update(ctx context.Context, user *model.User) error {
	return nil
}

func (r *UserRepository) Delete(ctx context.Context, id uint64) error {
	return nil
}


