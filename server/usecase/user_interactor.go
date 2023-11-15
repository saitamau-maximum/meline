package usecase

import (
	"context"

	"github.com/saitamau-maximum/meline/domain/entity"
	"github.com/saitamau-maximum/meline/domain/repository"
	model "github.com/saitamau-maximum/meline/models"
)

type IUserInteractor interface {
	GetUserByID(ctx context.Context, id uint64) (*entity.User, error)
	GetUserByGithubID(ctx context.Context, githubID string) (*entity.User, error)
	CreateUser(ctx context.Context, githubID, name, imageURL string) (*entity.User, error)
}

type UserInteractor struct {
	userRepository repository.IUserRepository
}

func NewUserInteractor(repository repository.IUserRepository) IUserInteractor {
	return &UserInteractor{
		userRepository: repository,
	}
}

func (i *UserInteractor) GetUserByID(ctx context.Context, id uint64) (*entity.User, error) {
	user, err := i.userRepository.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return user.ToUserEntity(), nil
}

func (i *UserInteractor) GetUserByGithubID(ctx context.Context, githubID string) (*entity.User, error) {
	user, err := i.userRepository.FindByGithubID(ctx, githubID)
	if err != nil {
		return nil, err
	}

	return user.ToUserEntity(), nil
}

func (i *UserInteractor) CreateUser(ctx context.Context, providerID, name, imageURL string) (*entity.User, error) {
	userModel := &model.User{
		ProviderID: providerID,
		Name: name,
		ImageURL: imageURL,
	}
	
	if err := i.userRepository.Create(ctx, userModel); err != nil {
		return nil, err
	}

	createdUser, err := i.userRepository.FindByGithubID(ctx, providerID)
	if err != nil {
		return nil, err
	}

	return createdUser.ToUserEntity(), nil
}
