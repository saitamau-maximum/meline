package usecase

import (
	"context"

	"github.com/saitamau-maximum/meline/domain/entity"
	"github.com/saitamau-maximum/meline/domain/repository"
)

type IGithubOAuthInteractor interface {
	GetGithubOAuthURL(ctx context.Context, state string) string
	GetGithubOAuthToken(ctx context.Context, code string) (string, error)
	GetGithubUser(ctx context.Context, token string) (*entity.OAuthUserResponse, error)
}

type GithubOAuthInteractor struct {
	oAuthRepository repository.IOAuthRepository
}

func NewGithubOAuthInteractor(r repository.IOAuthRepository) IGithubOAuthInteractor {
	return &GithubOAuthInteractor{
		oAuthRepository: r,
	}
}

func (i *GithubOAuthInteractor) GetGithubOAuthURL(ctx context.Context, state string) string {
	return i.oAuthRepository.GetOAuthURL(ctx, state)
}

func (i *GithubOAuthInteractor) GetGithubOAuthToken(ctx context.Context, code string) (string, error) {
	return i.oAuthRepository.GetOAuthToken(ctx, code)
}

func (i *GithubOAuthInteractor) GetGithubUser(ctx context.Context, token string) (*entity.OAuthUserResponse, error) {
	return i.oAuthRepository.GetUser(ctx, token)
}
