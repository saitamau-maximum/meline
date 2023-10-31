package usecase

import (
	"context"
	"io"
	"net/http"

	"github.com/saitamau-maximum/meline/domain/repository"
)

type IAuthInteractor interface {
	GetGithubOAuthURL(ctx context.Context, state string) string
	GetGithubOAuthToken(ctx context.Context, code string) (string, error)
	GetGithubUser(ctx context.Context, token string) (*http.Response, error)
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

func (i *AuthInteractor) GetGithubUser(ctx context.Context, token string) (*http.Response, error) {
	return i.repository.GetGithubUser(ctx, token)
}
