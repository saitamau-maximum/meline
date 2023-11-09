package auth

import (
	"context"
	"io"
	"net/http"
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"golang.org/x/oauth2"

	"github.com/saitamau-maximum/meline/domain/entity"
	"github.com/saitamau-maximum/meline/domain/repository"
)

const (
	GITHUB_ROOT_URL = "https://api.github.com/user"
)

type AuthRepository struct {
}

func NewAuthRepository() repository.IAuthRepository {
	return &AuthRepository{}
}

func (r *AuthRepository) GetGithubOAuthURL(ctx context.Context, state string) string {
	return NewGithubOAuthConf().AuthCodeURL(state)
}

func (r *AuthRepository) GetGithubOAuthToken(ctx context.Context, code string) (string, error) {
	token, err := NewGithubOAuthConf().Exchange(ctx, code)
	if err != nil {
		return "", err
	}

	return token.AccessToken, nil
}

func (r *AuthRepository) GetGithubUser(ctx context.Context, token string) ([]byte, error) {
	client := oauth2.NewClient(ctx, oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	))

	req, err := http.NewRequest("GET", GITHUB_ROOT_URL, nil)
	if err != nil {
		return nil, err
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return resBody, nil
}

func (r *AuthRepository) CreateAccessToken(ctx context.Context, user *entity.User) (string, error) {
	claims := jwt.MapClaims{
		"iss": "meline",
		"user_id": user.ID,
		"github_id": user.GithubID,
		"iat": time.Now().Unix(),
		"exp": time.Now().Add(time.Hour * 3).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}
