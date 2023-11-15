package usecase

import (
	"context"
	crand "crypto/rand"
	"errors"
	"fmt"
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"

	"github.com/saitamau-maximum/meline/domain/entity"
	"github.com/saitamau-maximum/meline/domain/repository"
)

type IOAuthInteractor interface {
	GetGithubOAuthURL(ctx context.Context, state string) string
	GetGithubOAuthToken(ctx context.Context, code string) (string, error)
	GetGithubUser(ctx context.Context, token string) (*entity.OAuthUserResponse, error)
	CreateAccessToken(ctx context.Context, user *entity.User) (string, error)
	GenerateState(b int) string
}

type OAuthInteractor struct {
	authRepository repository.IOAuthRepository
}

func NewOAuthInteractor(r repository.IOAuthRepository) IOAuthInteractor {
	return &OAuthInteractor{
		authRepository: r,
	}
}

func (i *OAuthInteractor) GetGithubOAuthURL(ctx context.Context, state string) string {
	return i.authRepository.GetOAuthURL(ctx, state)
}

func (i *OAuthInteractor) GetGithubOAuthToken(ctx context.Context, code string) (string, error) {
	return i.authRepository.GetOAuthToken(ctx, code)
}

func (i *OAuthInteractor) GetGithubUser(ctx context.Context, token string) (*entity.OAuthUserResponse, error) {
	return i.authRepository.GetUser(ctx, token)
}

func (i *OAuthInteractor) CreateAccessToken(ctx context.Context, user *entity.User) (string, error) {
	claims := jwt.MapClaims{
		"iss": "meline",
		"user_id": user.ID,
		"github_id": user.ProviderID,
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

func (i *OAuthInteractor) GenerateState(b int) string {
    k := make([]byte, b)
    if _, err := crand.Read(k); err != nil {
        panic(err)
    }
    return fmt.Sprintf("%x", k)
}
