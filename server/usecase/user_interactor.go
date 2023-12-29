package usecase

import (
	"context"
	"database/sql"

	"github.com/saitamau-maximum/meline/domain/repository"
	"github.com/saitamau-maximum/meline/models"
	"github.com/saitamau-maximum/meline/usecase/presenter"
)

type IUserInteractor interface {
	GetUserByID(ctx context.Context, id uint64) (*presenter.GetUserByIdResponse, error)
	GetUserByGithubID(ctx context.Context, githubID, userName, imageUrl string) (*presenter.GetUserByGithubIdResponse, error)
	CreateUser(ctx context.Context, githubID, name, imageURL string) (*presenter.CreateUserResponse, error)
}

type UserInteractor struct {
	userRepository repository.IUserRepository
	userPresenter  presenter.IUserPresenter
}

func NewUserInteractor(repository repository.IUserRepository, userPresenter presenter.IUserPresenter) IUserInteractor {
	return &UserInteractor{
		userRepository: repository,
		userPresenter:  userPresenter,
	}
}

func (i *UserInteractor) GetUserByID(ctx context.Context, id uint64) (*presenter.GetUserByIdResponse, error) {
	user, err := i.userRepository.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return i.userPresenter.GenerateGetUserByIdResponse(user.ToUserEntity()), nil
}

func (i *UserInteractor) GetUserByGithubID(ctx context.Context, githubID, userName, imageUrl string) (*presenter.GetUserByGithubIdResponse, error) {
	user, err := i.userRepository.FindByProviderID(ctx, githubID)
	if err != nil {
		if err := sql.ErrNoRows; err != nil {
			newUser := &model.User{
				ProviderID: githubID,
				Name:       userName,
				ImageURL:   imageUrl,
			}
	
			if _err := i.userRepository.Create(ctx, newUser); _err != nil {
				return &presenter.GetUserByGithubIdResponse{}, _err
			}
	
			createdUser, _err := i.userRepository.FindByProviderID(ctx, githubID)
			if _err != nil {
				return &presenter.GetUserByGithubIdResponse{}, _err
			}
	
			user = createdUser
		} else {
			return &presenter.GetUserByGithubIdResponse{}, err
		}
	}

	return i.userPresenter.GenerateGetUserByGithubIdResponse(user.ToUserEntity()), nil
}

func (i *UserInteractor) CreateUser(ctx context.Context, providerID, name, imageURL string) (*presenter.CreateUserResponse, error) {
	userModel := &model.User{
		ProviderID: providerID,
		Name:       name,
		ImageURL:   imageURL,
	}

	if err := i.userRepository.Create(ctx, userModel); err != nil {
		return nil, err
	}

	createdUser, err := i.userRepository.FindByProviderID(ctx, providerID)
	if err != nil {
		return nil, err
	}

	return i.userPresenter.GenerateCreateUserResponse(createdUser.ToUserEntity()), nil
}
