package presenter

import (
	"github.com/saitamau-maximum/meline/domain/entity"
)

type UserMeResponse struct {
	ID         uint64 `json:"id"`
	Name	   string `json:"name"`
	ImageURL   string `json:"image_url"`
}

type IUserPresenter interface {
	GenreateUserMeResponse(user *entity.User) *UserMeResponse
}
