package usecase_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/saitamau-maximum/meline/domain/entity"
	"github.com/saitamau-maximum/meline/models"
	"github.com/saitamau-maximum/meline/usecase"
	"github.com/saitamau-maximum/meline/usecase/presenter"

	"github.com/stretchr/testify/assert"
)

func TestUserInteractor_Success_GetUserByID(t *testing.T) {
	ctx := context.Background()
	repo := &mockUserRepository{}
	pre := &mockUserPresenter{}

	interactor := usecase.NewUserInteractor(repo, pre)

	expectedUser := &presenter.GetUserByIdResponse{
		ID:       1,
		Name:     "John Doe",
		ImageURL: "https://example.com/image.jpg",
	}

	result, err := interactor.GetUserByID(ctx, 1)
	assert.NoError(t, err)
	assert.Equal(t, expectedUser, result)
}

func TestUserInteractor_Failed_GetUserByID_NotFound(t *testing.T) {
	ctx := context.Background()
	repo := &mockUserRepository{}
	pre := &mockUserPresenter{}

	interactor := usecase.NewUserInteractor(repo, pre)
	ctx = context.WithValue(ctx, FindFailedValue, true)

	result, err := interactor.GetUserByID(ctx, 2)
	assert.Error(t, err)
	assert.Nil(t, result)
}

func TestUserInteractor_Success_GetUserByGithubID(t *testing.T) {
	ctx := context.Background()
	repo := &mockUserRepository{}
	pre := &mockUserPresenter{}

	interactor := usecase.NewUserInteractor(repo, pre)

	expectedUser := &presenter.GetUserByGithubIdResponse{
		ID:       1,
		Name:     "John Doe",
		ImageURL: "https://example.com/image.jpg",
	}

	result, err := interactor.GetUserByGithubIDOrCreate(ctx, "test-provider-id", "John Doe", "https://example.com/image.jpg")
	assert.NoError(t, err)
	assert.Equal(t, expectedUser, result)
}

func TestUserInteractor_Failed_GetUserByGithubID_NotFound(t *testing.T) {
	ctx := context.Background()
	repo := &mockUserRepository{}
	pre := &mockUserPresenter{}

	interactor := usecase.NewUserInteractor(repo, pre)

	ctx = context.WithValue(ctx, FindByProviderIDFailedValue, true)

	result, err := interactor.GetUserByGithubIDOrCreate(ctx, "test-provider-id", "John Doe", "https://example.com/image.jpg")

	expectedUser := &presenter.GetUserByGithubIdResponse{}

	assert.Error(t, err)
	assert.Equal(t, expectedUser, result)
}

func TestUserInteractor_Success_CreateUser(t *testing.T) {
	ctx := context.Background()
	repo := &mockUserRepository{}
	pre := &mockUserPresenter{}

	interactor := usecase.NewUserInteractor(repo, pre)

	expectedUser := &presenter.CreateUserResponse{
		ID: 1,
	}

	user, err := interactor.CreateUser(ctx, "test-provider-id", "John Doe", "https://example.com/image.jpg")
	assert.NoError(t, err)
	assert.Equal(t, expectedUser, user)
}

func TestUserInteractor_Failed_CreateUser_CreateFailed(t *testing.T) {
	ctx := context.Background()
	repo := &mockUserRepository{}
	pre := &mockUserPresenter{}

	interactor := usecase.NewUserInteractor(repo, pre)
	ctx = context.WithValue(ctx, CreateFailedValue, true)

	user, err := interactor.CreateUser(ctx, "test-provider-id", "John Doe", "https://example.com/image.jpg")
	assert.Error(t, err)
	assert.Nil(t, user)
}

func TestUserInteractor_Failed_CreateUser_FindFailed(t *testing.T) {
	ctx := context.Background()
	repo := &mockUserRepository{}
	pre := &mockUserPresenter{}

	interactor := usecase.NewUserInteractor(repo, pre)
	ctx = context.WithValue(ctx, FindByProviderIDFailedValue, true)

	user, err := interactor.CreateUser(ctx, "test-provider-id", "John Doe", "https://example.com/image.jpg")
	assert.Error(t, err)
	assert.Nil(t, user)
}

type mockUserRepository struct{}

type FindFailed string
type FindByProviderIDFailed string
type CreateFailed string
type FindChannelsFailed string

const (
	FindFailedValue             FindFailed             = "find_failed"
	FindByProviderIDFailedValue FindByProviderIDFailed = "find_by_provider_id_failed"
	CreateFailedValue           CreateFailed           = "create_failed"
	FindChannelsFailedValue     FindChannelsFailed     = "find_channel_failed"
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

func (r *mockUserRepository) FindChannelsByUserID(ctx context.Context, userID uint64) ([]*model.Channel, error) {
	if ctx.Value(FindChannelsFailedValue) != nil {
		return nil, fmt.Errorf("failed to find channels")
	}

	return []*model.Channel{
		{
			ID:   1,
			Name: "test-channel",
		},
	}, nil
}

type mockUserPresenter struct{}

func (p *mockUserPresenter) GenerateGetUserByIdResponse(user *entity.User) *presenter.GetUserByIdResponse {
	return &presenter.GetUserByIdResponse{
		ID:       user.ID,
		Name:     user.Name,
		ImageURL: user.ImageURL,
	}
}

func (p *mockUserPresenter) GenerateGetUserByGithubIdResponse(user *entity.User) *presenter.GetUserByGithubIdResponse {
	return &presenter.GetUserByGithubIdResponse{
		ID:       user.ID,
		Name:     user.Name,
		ImageURL: user.ImageURL,
	}
}

func (p *mockUserPresenter) GenerateCreateUserResponse(user *entity.User) *presenter.CreateUserResponse {
	return &presenter.CreateUserResponse{
		ID: user.ID,
	}
}
