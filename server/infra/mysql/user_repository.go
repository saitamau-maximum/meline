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

func (r *UserRepository) FindByID(ctx context.Context, id uint64) (*model.User, error) {
	var user model.User

	if err := r.db.NewSelect().Model(&model.User{}).Where("id = ?", id).Where("deleted_at IS NULL").Scan(ctx, &user); err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) FindByGithubID(ctx context.Context, githubID string) (*model.User, error) {
	var user model.User

	if err := r.db.NewSelect().Model(&model.User{}).Where("github_id = ?", githubID).Where("deleted_at IS NULL").Scan(ctx, &user); err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) Insert(ctx context.Context, githubID, name string) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	user := model.User{
		GithubID: githubID,
		Name:     name,
	}

	if _, err := tx.NewInsert().Model(&user).Exec(ctx); err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}

func (r *UserRepository) Update(ctx context.Context, id uint64, githubID, name string) error {
	return nil
}

func (r *UserRepository) Delete(ctx context.Context, id uint64) error {
	return nil
}


