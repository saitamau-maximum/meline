package repository

import (
	"context"

	"github.com/saitamau-maximum/meline/domain/entity"
)

type IAuthRepository interface {
	GetOAuthURL(ctx context.Context, state string) string
	GetOAuthToken(ctx context.Context, code string) (string, error)
	GetUser(ctx context.Context, token string) (*entity.OAuthUserResponse, error)
}
