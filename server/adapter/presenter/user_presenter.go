package presenter

import (
	"github.com/saitamau-maximum/meline/domain/entity"
	"github.com/saitamau-maximum/meline/usecase/presenter"
)

type UserPresenter struct{}

func NewUserPresenter() presenter.IUserPresenter {
	return &UserPresenter{}
}

func (p *UserPresenter) GenreateUserMeResponse(user *entity.User) *presenter.UserMeResponse {
	return &presenter.UserMeResponse{
		ID:       user.ID,
		Name:     user.Name,
		ImageURL: user.ImageURL,
	}
}
