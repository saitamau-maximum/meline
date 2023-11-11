package usecase

import (
	"context"
	"errors"
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"

	"github.com/saitamau-maximum/meline/domain/entity"
	"github.com/saitamau-maximum/meline/domain/repository"
)

type IAuthInteractor interface {
	GetGithubOAuthURL(ctx context.Context, state string) string
	GetGithubOAuthToken(ctx context.Context, code string) (string, error)
	GetGithubUser(ctx context.Context, token string) (map[string]interface{}, error)
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

func (i *AuthInteractor) GetGithubUser(ctx context.Context, token string) (map[string]interface{}, error) {
	return i.repository.GetGithubUser(ctx, token)
}

func (i *AuthInteractor) CreateAccessToken(ctx context.Context, user *entity.User) (string, error) {
	claims := jwt.MapClaims{
		"iss": "meline",
		"user_id": user.ID,
		"github_id": user.GithubID,
		"iat": time.Now().Unix(),
		"exp": time.Now().Add(time.Hour * 3).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		return "", errors.New("JWT_SECRET is not set")
	}

	return token.SignedString([]byte(jwtSecret))
}
