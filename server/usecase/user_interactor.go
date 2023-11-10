package usecase

import (
	"context"

	"github.com/saitamau-maximum/meline/domain/entity"
	"github.com/saitamau-maximum/meline/domain/repository"
)

type IUserInteractor interface {
	GetUserByID(ctx context.Context, id uint64) (*entity.User, error)
	GetUserByGithubID(ctx context.Context, githubID string) (*entity.User, error)
	CreateUser(ctx context.Context, user *entity.User) error
}

type UserInteractor struct {
	repository repository.IUserRepository
}

func NewUserInteractor(repository repository.IUserRepository) IUserInteractor {
	return &UserInteractor{
		repository: repository,
	}
}

func (i *UserInteractor) GetUserByID(ctx context.Context, id uint64) (*entity.User, error) {
	user, err := i.repository.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return user.ToUserEntity(), nil
}

func (i *UserInteractor) GetUserByGithubID(ctx context.Context, githubID string) (*entity.User, error) {
	user, err := i.repository.FindByGithubID(ctx, githubID)
	if err != nil {
		return nil, err
	}

	return user.ToUserEntity(), nil
}

func (i *UserInteractor) CreateUser(ctx context.Context, user *entity.User) error {
	return i.repository.Create(ctx, user.GithubID, user.Name)
}
