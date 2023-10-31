package auth

import (
	"context"
	"net/http"

	"golang.org/x/oauth2"

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

func (r *AuthRepository) GetGithubUser(ctx context.Context, token string) (*http.Response, error) {
	client := oauth2.NewClient(ctx, oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	))

	req, err := http.NewRequest("GET", GITHUB_ROOT_URL, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", "application/vnd.github.v3+json")

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
