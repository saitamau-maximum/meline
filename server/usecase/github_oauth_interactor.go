package usecase

import (
	"context"
	"time"

	jwt "github.com/dgrijalva/jwt-go"

	"github.com/saitamau-maximum/meline/config"
	"github.com/saitamau-maximum/meline/domain/entity"
	"github.com/saitamau-maximum/meline/domain/repository"
)

type IGithubOAuthInteractor interface {
	GetGithubOAuthURL(ctx context.Context, state string) string
	GetGithubOAuthToken(ctx context.Context, code string) (string, error)
	GetGithubUser(ctx context.Context, token string) (*entity.OAuthUserResponse, error)
	CreateAccessToken(ctx context.Context, user *entity.User) (string, error)
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

func (i *GithubOAuthInteractor) CreateAccessToken(ctx context.Context, user *entity.User) (string, error) {
	claims := jwt.MapClaims{
		"iss":         config.APP_IDENTIFIER,
		"user_id":     user.ID,
		"provider_id": user.ProviderID,
		"iat":         time.Now().Unix(),
		"exp":         time.Now().Add(config.ACCESS_TOKEN_EXPIRE).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(config.JWT_SECRET))
}
