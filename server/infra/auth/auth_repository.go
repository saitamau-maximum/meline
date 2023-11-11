package auth

import (
	"context"
	"encoding/json"
	"io"
	"net/http"

	"golang.org/x/oauth2"

	"github.com/saitamau-maximum/meline/domain/repository"
)

const (
	GITHUB_ROOT_URL = "https://api.github.com/user"
)

type AuthRepository struct {
	OAuthConf *oauth2.Config
}

func NewAuthRepository(conf *oauth2.Config) repository.IAuthRepository {
	return &AuthRepository{
		OAuthConf: conf,
	}
}

func (r *AuthRepository) GetGithubOAuthURL(ctx context.Context, state string) string {
	return r.OAuthConf.AuthCodeURL(state)
}

func (r *AuthRepository) GetGithubOAuthToken(ctx context.Context, code string) (string, error) {
	token, err := r.OAuthConf.Exchange(ctx, code)
	if err != nil {
		return "", err
	}

	return token.AccessToken, nil
}

func (r *AuthRepository) GetGithubUser(ctx context.Context, token string) (map[string]interface{}, error) {
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

	var gitRes map[string]interface{}
	if err := json.Unmarshal(resBody, &gitRes); err != nil {
		return nil, err
	}

	return gitRes, nil
}

