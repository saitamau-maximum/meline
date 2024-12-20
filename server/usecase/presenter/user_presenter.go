package presenter

import (
	"github.com/saitamau-maximum/meline/domain/entity"
)

type User struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	ImageURL string `json:"image_url"`
}

type GetUserByIdResponse struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	ImageURL string `json:"image_url"`
}

type GetUserByGithubIdResponse struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	ImageURL string `json:"image_url"`
}

type CreateUserResponse struct {
	ID string `json:"id"`
}

type IUserPresenter interface {
	GenerateGetUserByIdResponse(user *entity.User) *GetUserByIdResponse
	GenerateGetUserByGithubIdResponse(user *entity.User) *GetUserByGithubIdResponse
	GenerateCreateUserResponse(user *entity.User) *CreateUserResponse
}
