package usecase

import (
	"context"

	"github.com/saitamau-maximum/meline/domain/entity"
	"github.com/saitamau-maximum/meline/domain/repository"
)

type IUserInteractor interface {
	GetUser(ctx context.Context, id uint64) (*entity.User, error)
	GetUserByGithubID(ctx context.Context, githubID string) (*entity.User, error)
	CreateUser(ctx context.Context, githubID, name string) error
	UpdateUser(ctx context.Context, id uint64, githubID, name string) error
	DeleteUser(ctx context.Context, id uint64) error
}

type UserInteractor struct {
	repository repository.IUserRepository
}

func NewUserInteractor(repository repository.IUserRepository) IUserInteractor {
	return &UserInteractor{
		repository: repository,
	}
}

func (i *UserInteractor) GetUser(ctx context.Context, id uint64) (*entity.User, error) {
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

func (i *UserInteractor) CreateUser(ctx context.Context, githubID, name string) error {
	if err := i.repository.Insert(ctx, githubID, name); err != nil {
		return err
	}

	return nil
}

func (i *UserInteractor) UpdateUser(ctx context.Context, id uint64, githubID, name string) error {
	if err := i.repository.Update(ctx, id, githubID, name); err != nil {
		return err
	}

	return nil
}

func (i *UserInteractor) DeleteUser(ctx context.Context, id uint64) error {
	if err := i.repository.Delete(ctx, id); err != nil {
		return err
	}

	return nil
}
