package usecase_test

import (
	"context"
	"testing"

	"github.com/saitamau-maximum/meline/domain/entity"
	"github.com/saitamau-maximum/meline/usecase"
	"github.com/stretchr/testify/assert"
)

func TestGithubOAuthInteractor_GetGithubOAuthURL(t *testing.T) {
	repo := &mockOAuthRepository{}

	interactor := usecase.NewGithubOAuthInteractor(repo)

	expectedURL := "https://github.com/oauth/auth"

	result := interactor.GetGithubOAuthURL(context.Background(), "state")
	assert.Equal(t, expectedURL, result)
}

func TestGithubOAuthInteractor_GetGithubOAuthToken(t *testing.T) {
	repo := &mockOAuthRepository{}

	interactor := usecase.NewGithubOAuthInteractor(repo)

	expectedToken := "test-token"

	result, err := interactor.GetGithubOAuthToken(context.Background(), "code")
	assert.NoError(t, err)
	assert.Equal(t, expectedToken, result)
}

func TestGithubOAuthInteractor_GetGithubUser(t *testing.T) {
	repo := &mockOAuthRepository{}

	interactor := usecase.NewGithubOAuthInteractor(repo)

	expectedUser := &entity.OAuthUserResponse{
		Name:     "John Doe",
		ImageURL: "https://example.com/image.jpg",
	}

	result, err := interactor.GetGithubUser(context.Background(), "token")
	assert.NoError(t, err)
	assert.Equal(t, expectedUser, result)
}

type mockOAuthRepository struct{}

func (r *mockOAuthRepository) GetOAuthURL(ctx context.Context, state string) string {
	return "https://github.com/oauth/auth"
}

func (r *mockOAuthRepository) GetOAuthToken(ctx context.Context, code string) (string, error) {
	return "test-token", nil
}

func (r *mockOAuthRepository) GetUser(ctx context.Context, token string) (*entity.OAuthUserResponse, error) {
	return &entity.OAuthUserResponse{
		Name:     "John Doe",
		ImageURL: "https://example.com/image.jpg",
	}, nil
}
