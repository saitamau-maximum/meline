package github

import (
	"context"
	"encoding/json"
	"io"
	"net/http"

	"golang.org/x/oauth2"

	"github.com/saitamau-maximum/meline/domain/entity"
	"github.com/saitamau-maximum/meline/domain/repository"
)

const (
	GITHUB_ROOT_URL = "https://api.github.com/user"
)

type OAuthRepository struct {
	OAuthConf *oauth2.Config
}

func NewOAuthRepository(conf *oauth2.Config) repository.IOAuthRepository {
	return &OAuthRepository{
		OAuthConf: conf,
	}
}

func (r *OAuthRepository) GetOAuthURL(ctx context.Context, state string) string {
	return r.OAuthConf.AuthCodeURL(state)
}

func (r *OAuthRepository) GetOAuthToken(ctx context.Context, code string) (string, error) {
	token, err := r.OAuthConf.Exchange(ctx, code)
	if err != nil {
		return "", err
	}

	return token.AccessToken, nil
}

func (r *OAuthRepository) GetUser(ctx context.Context, token string) (*entity.OAuthUserResponse, error) {
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

	return &entity.OAuthUserResponse{
		OAuthUserID:   gitRes["login"].(string),
		Name:     	   gitRes["name"].(string),
		ImageURL:	   gitRes["avatar_url"].(string),
	}, nil
}

