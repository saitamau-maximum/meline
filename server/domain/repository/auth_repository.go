package repository

import (
	"context"
	"net/http"
)

type IAuthRepository interface {
	GetGithubOAuthURL(ctx context.Context, state string) string
	GetGithubOAuthToken(ctx context.Context, code string) (string, error)
	GetGithubUser(ctx context.Context, token string) (*http.Response, error)
}
