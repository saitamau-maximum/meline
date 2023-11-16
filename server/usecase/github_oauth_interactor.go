package usecase

import (
	"context"
	crand "crypto/rand"
	"errors"
	"fmt"
	"net/http"
	"os"
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
	GenerateState(stateLength int) string
	GenerateStateCookie(state string, isDev bool) *http.Cookie
	GenerateAccessTokenCookie(token string, isDev bool) *http.Cookie
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
		"iss":         "meline",
		"user_id":     user.ID,
		"provider_id": user.ProviderID,
		"iat":         time.Now().Unix(),
		"exp":         time.Now().Add(time.Hour * 24 * 7).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		return "", errors.New("JWT_SECRET is not set")
	}

	return token.SignedString([]byte(jwtSecret))
}

func (i *GithubOAuthInteractor) GenerateState(stateLength int) string {
	k := make([]byte, stateLength)
	if _, err := crand.Read(k); err != nil {
		panic(err)
	}
	return fmt.Sprintf("%x", k)
}

func (i *GithubOAuthInteractor) GenerateStateCookie(state string, isDev bool) *http.Cookie {
	cookie := new(http.Cookie)
	cookie.Name = config.OAUTH_STATE_COOKIE_NAME
	cookie.Value = state
	cookie.Path = "/"
	cookie.HttpOnly = true
	cookie.SameSite = http.SameSiteLaxMode
	cookie.Secure = !isDev
	cookie.Expires = time.Now().Add(5 * time.Minute)

	return cookie
}

func (i *GithubOAuthInteractor) GenerateAccessTokenCookie(token string, isDev bool) *http.Cookie {
	cookie := new(http.Cookie)
	cookie.Name = config.ACCESS_TOKEN_COOKIE_NAME
	cookie.Value = token
	cookie.Path = "/"
	cookie.HttpOnly = true
	cookie.SameSite = http.SameSiteLaxMode
	cookie.Secure = !isDev
	cookie.Expires = time.Now().Add(3 * time.Hour)

	return cookie
}
