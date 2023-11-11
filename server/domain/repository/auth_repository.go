package repository

import (
	"context"
)

type IAuthRepository interface {
	GetGithubOAuthURL(ctx context.Context, state string) string
	GetGithubOAuthToken(ctx context.Context, code string) (string, error)
	GetGithubUser(ctx context.Context, token string) (map[string]interface{}, error)
}
