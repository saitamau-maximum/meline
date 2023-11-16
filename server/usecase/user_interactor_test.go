package usecase_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/saitamau-maximum/meline/domain/entity"
	model "github.com/saitamau-maximum/meline/models"
	"github.com/saitamau-maximum/meline/usecase"

	"github.com/stretchr/testify/assert"
)

func TestUserInteractor_Success_GetUserByID(t *testing.T) {
	ctx := context.Background()
	repo := &mockUserRepository{}

	interactor := usecase.NewUserInteractor(repo)

	expectedUser := &entity.User{
		ID:         1,
		ProviderID: "test-provider-id",
		Name:       "John Doe",
		ImageURL:   "https://example.com/image.jpg",
	}

	result, err := interactor.GetUserByID(ctx, 1)
	assert.NoError(t, err)
	assert.Equal(t, expectedUser, result)
}

func TestUserInteractor_Failed_GetUserByID_NotFound(t *testing.T) {
	ctx := context.Background()
	repo := &mockUserRepository{}

	interactor := usecase.NewUserInteractor(repo)
	ctx = context.WithValue(ctx, FindFailedValue, true)

	result, err := interactor.GetUserByID(ctx, 2)
	assert.Error(t, err)
	assert.Nil(t, result)
}

func TestUserInteractor_Success_GetUserByGithubID(t *testing.T) {
	ctx := context.Background()
	repo := &mockUserRepository{}

	interactor := usecase.NewUserInteractor(repo)

	expectedUser := &entity.User{
		ID:         1,
		ProviderID: "test-provider-id",
		Name:       "John Doe",
		ImageURL:   "https://example.com/image.jpg",
	}

	result, err := interactor.GetUserByGithubID(ctx, "test-provider-id")
	assert.NoError(t, err)
	assert.Equal(t, expectedUser, result)
}

func TestUserInteractor_Failed_GetUserByGithubID_NotFound(t *testing.T) {
	ctx := context.Background()
	repo := &mockUserRepository{}

	interactor := usecase.NewUserInteractor(repo)

	ctx = context.WithValue(ctx, FindByProviderIDFailedValue, true)

	result, err := interactor.GetUserByGithubID(ctx, "test-provider-id")
	assert.Error(t, err)
	assert.Nil(t, result)
}

func TestUserInteractor_Success_CreateUser(t *testing.T) {
	ctx := context.Background()
	repo := &mockUserRepository{}

	interactor := usecase.NewUserInteractor(repo)

	user, err := interactor.CreateUser(ctx, "test-provider-id", "John Doe", "https://example.com/image.jpg")
	assert.NoError(t, err)
	assert.Equal(t, uint64(1), user.ID)
}

func TestUserInteractor_Failed_CreateUser_CreateFailed(t *testing.T) {
	ctx := context.Background()
	repo := &mockUserRepository{}

	interactor := usecase.NewUserInteractor(repo)
	ctx = context.WithValue(ctx, CreateFailedValue, true)

	user, err := interactor.CreateUser(ctx, "test-provider-id", "John Doe", "https://example.com/image.jpg")
	assert.Error(t, err)
	assert.Nil(t, user)
}

func TestUserInteractor_Failed_CreateUser_FindFailed(t *testing.T) {
	ctx := context.Background()
	repo := &mockUserRepository{}

	interactor := usecase.NewUserInteractor(repo)
	ctx = context.WithValue(ctx, FindByProviderIDFailedValue, true)

	user, err := interactor.CreateUser(ctx, "test-provider-id", "John Doe", "https://example.com/image.jpg")
	assert.Error(t, err)
	assert.Nil(t, user)
}

type mockUserRepository struct{}

type FindFailed string
type FindByProviderIDFailed string
type CreateFailed string

const (
	FindFailedValue             FindFailed             = "find_failed"
	FindByProviderIDFailedValue FindByProviderIDFailed = "find_by_provider_id_failed"
	CreateFailedValue           CreateFailed           = "create_failed"
)

func (r *mockUserRepository) FindByID(ctx context.Context, id uint64) (*model.User, error) {
	if ctx.Value(FindFailedValue) != nil {
		return nil, fmt.Errorf("not found")
	}

	return &model.User{
		ID:         1,
		ProviderID: "test-provider-id",
		Name:       "John Doe",
		ImageURL:   "https://example.com/image.jpg",
	}, nil
}

func (r *mockUserRepository) FindByProviderID(ctx context.Context, providerID string) (*model.User, error) {
	if ctx.Value(FindByProviderIDFailedValue) != nil {
		return nil, fmt.Errorf("not found")
	}

	return &model.User{
		ID:         1,
		ProviderID: "test-provider-id",
		Name:       "John Doe",
		ImageURL:   "https://example.com/image.jpg",
	}, nil
}

func (r *mockUserRepository) Create(ctx context.Context, user *model.User) error {
	if ctx.Value(CreateFailedValue) != nil {
		return fmt.Errorf("failed to create user")
	}

	return nil
}
