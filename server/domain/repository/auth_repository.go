package repository

import (
	"context"

	"github.com/saitamau-maximum/meline/domain/entity"
)

type IAuthRepository interface {
	GetGithubOAuthURL(ctx context.Context, state string) string
	GetGithubOAuthToken(ctx context.Context, code string) (string, error)
	GetGithubUser(ctx context.Context, token string) ([]byte, error)
	CreateAccessToken(ctx context.Context, user *entity.User) (string, error)
}
