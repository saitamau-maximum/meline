package usecase

import (
	"context"

	"github.com/saitamau-maximum/meline/domain/entity"
	"github.com/saitamau-maximum/meline/domain/repository"
)

type IAuthInteractor interface {
	GetGithubOAuthURL(ctx context.Context, state string) string
	GetGithubOAuthToken(ctx context.Context, code string) (string, error)
	GetGithubUser(ctx context.Context, gitToken string) ([]byte, error)
	CreateAccessToken(ctx context.Context, user *entity.User) (string, error)
}

type AuthInteractor struct {
	repository repository.IAuthRepository
}

func NewAuthInteractor(r repository.IAuthRepository) IAuthInteractor {
	return &AuthInteractor{
		repository: r,
	}
}

func (i *AuthInteractor) GetGithubOAuthURL(ctx context.Context, state string) string {
	return i.repository.GetGithubOAuthURL(ctx, state)
}

func (i *AuthInteractor) GetGithubOAuthToken(ctx context.Context, code string) (string, error) {
	return i.repository.GetGithubOAuthToken(ctx, code)
}

func (i *AuthInteractor) GetGithubUser(ctx context.Context, gitToken string) ([]byte, error) {
	return i.repository.GetGithubUser(ctx, gitToken)
}

func (i *AuthInteractor) CreateAccessToken(ctx context.Context, user *entity.User) (string, error) {
	return i.repository.CreateAccessToken(ctx, user)
}
