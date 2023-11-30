package presenter

import (
	"github.com/saitamau-maximum/meline/domain/entity"
	"github.com/saitamau-maximum/meline/usecase/presenter"
)

type UserPresenter struct{}

func NewUserPresenter() presenter.IUserPresenter {
	return &UserPresenter{}
}

func (p *UserPresenter) GenerateGetUserByIdResponse(user *entity.User) *presenter.GetUserByIdResponse {
	return &presenter.GetUserByIdResponse{
		ID:       user.ID,
		Name:     user.Name,
		ImageURL: user.ImageURL,
	}
}

func (p *UserPresenter) GenerateGetUserByGithubIdResponse(user *entity.User) *presenter.GetUserByGithubIdResponse {
	return &presenter.GetUserByGithubIdResponse{
		ID:       user.ID,
		Name:     user.Name,
		ImageURL: user.ImageURL,
	}
}

func (p *UserPresenter) GenerateCreateUserResponse(user *entity.User) *presenter.CreateUserResponse {
	return &presenter.CreateUserResponse{
		ID: user.ID,
	}
}
