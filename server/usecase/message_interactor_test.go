package usecase_test

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/saitamau-maximum/meline/config"
	"github.com/saitamau-maximum/meline/domain/entity"
	model "github.com/saitamau-maximum/meline/models"
	"github.com/saitamau-maximum/meline/usecase"
	"github.com/saitamau-maximum/meline/usecase/presenter"

	"github.com/stretchr/testify/assert"
)

type GetMessagesByChannelIDFailed string
type CreateMessageFailed string
type CreateReplyFailed string
type UpdateMessageFailed string
type DeleteMessageFailed string

type ExpectedCreateMessageResponse struct {
	msg    *presenter.CreateMessageResponse
	notify *presenter.NotifyMessageResponse
}

type ExpectedCreateReplyResponse struct {
	msg    *presenter.CreateMessageResponse
	notify *presenter.NotifyMessageResponse
}

const (
	GetMessagesByChannelIDFailedValue GetMessagesByChannelIDFailed = "get_messages_by_channel_id_failed"
	CreateMessageFailedValue          CreateMessageFailed          = "create_message_failed"
	CreateReplyFailedValue            CreateReplyFailed            = "create_reply_failed"
	UpdateMessageFailedValue          UpdateMessageFailed          = "update_message_failed"
	DeleteMessageFailedValue          DeleteMessageFailed          = "delete_message_failed"
)

func TestMessageInteractor_Success_GetMessagesByChannelID(t *testing.T) {
	ctx := context.Background()

	repo := &mockMessageRepository{}
	userRepo := &mockUserRepository{}
	notifyRepo := &mockNotifyRepository{}
	msgPre := &mockMessagePresenter{}
	notifyPre := &mockNotifyPresenter{}
	interactor := usecase.NewMessageInteractor(repo, userRepo, notifyRepo, msgPre, notifyPre)

	res, err := interactor.GetMessagesByChannelID(ctx, 1)

	expected := &presenter.GetMessagesByChannelIDResponse{
		Messages: []*presenter.Message{
			{
				ID: "1",
				User: &presenter.User{
					ID:       "1",
					Name:     "User",
					ImageURL: "https://example.com/image.png",
				},
				Content:        "Hello, World!",
				ReplyToMessage: nil,
				CreatedAt:      "0001-01-01 00:00:00 +0000 UTC",
				UpdatedAt:      "0001-01-01 00:00:00 +0000 UTC",
			},
		},
	}

	assert.NoError(t, err)
	assert.Equal(t, expected, res)
}

func TestMessageInteractor_Success_Create(t *testing.T) {
	ctx := context.Background()

	repo := &mockMessageRepository{}
	userRepo := &mockUserRepository{}
	notifyRepo := &mockNotifyRepository{}
	msgPre := &mockMessagePresenter{}
	notifyPre := &mockNotifyPresenter{}
	interactor := usecase.NewMessageInteractor(repo, userRepo, notifyRepo, msgPre, notifyPre)

	res, notifyRes, userIDs, err := interactor.Create(ctx, 1, 1, "Hello, World!")

	expected := &ExpectedCreateMessageResponse{
		msg: &presenter.CreateMessageResponse{
			Message: &presenter.Message{
				ID: "1",
				User: &presenter.User{
					ID:       "1",
					Name:     "User",
					ImageURL: "https://example.com/image.png",
				},
				Content:        "Hello, World!",
				ReplyToMessage: nil,
				CreatedAt:      "0001-01-01 00:00:00 +0000 UTC",
				UpdatedAt:      "0001-01-01 00:00:00 +0000 UTC",
			},
			ChannelID: "1",
		},
		notify: &presenter.NotifyMessageResponse{
			NotifyMeta: presenter.NotifyMeta{
				TypeID: config.NOTIFY_MESSAGE,
			},
			Payload: presenter.Payload{
				Message: &presenter.Message{
					ID: "1",
					User: &presenter.User{
						ID:       "1",
						Name:     "User",
						ImageURL: "https://example.com/image.png",
					},
					Content:        "Hello, World!",
					ReplyToMessage: nil,
					CreatedAt:      "0001-01-01 00:00:00 +0000 UTC",
					UpdatedAt:      "0001-01-01 00:00:00 +0000 UTC",
				},
				ChannelID: "1",
			},
		},
	}

	assert.NoError(t, err)
	assert.Equal(t, expected.msg, res)
	assert.Equal(t, expected.notify, notifyRes)
	assert.Equal(t, []uint64{1, 2}, userIDs)
}

func TestMessageInteractor_Success_CreateReply(t *testing.T) {
	ctx := context.Background()

	repo := &mockMessageRepository{}
	userRepo := &mockUserRepository{}
	notifyRepo := &mockNotifyRepository{}
	msgPre := &mockMessagePresenter{}
	notifyPre := &mockNotifyPresenter{}
	interactor := usecase.NewMessageInteractor(repo, userRepo, notifyRepo, msgPre, notifyPre)

	res, notifyRes, userIDs, err := interactor.CreateReply(ctx, 1, 1, "1", "Hello, World!")

	expected := &ExpectedCreateReplyResponse{
		msg: &presenter.CreateMessageResponse{
			Message: &presenter.Message{
				ID: "1",
				User: &presenter.User{
					ID:       "1",
					Name:     "User",
					ImageURL: "https://example.com/image.png",
				},
				Content:        "Hello, World!",
				ReplyToMessage: nil,
				CreatedAt:      "0001-01-01 00:00:00 +0000 UTC",
				UpdatedAt:      "0001-01-01 00:00:00 +0000 UTC",
			},
			ChannelID: "1",
		},
		notify: &presenter.NotifyMessageResponse{
			NotifyMeta: presenter.NotifyMeta{
				TypeID: config.NOTIFY_MESSAGE,
			},
			Payload: presenter.Payload{
				Message: &presenter.Message{
					ID: "1",
					User: &presenter.User{
						ID:       "1",
						Name:     "User",
						ImageURL: "https://example.com/image.png",
					},
					Content:        "Hello, World!",
					ReplyToMessage: nil,
					CreatedAt:      "0001-01-01 00:00:00 +0000 UTC",
					UpdatedAt:      "0001-01-01 00:00:00 +0000 UTC",
				},
				ChannelID: "1",
			},
		},
	}

	assert.NoError(t, err)
	assert.Equal(t, expected.msg, res)
	assert.Equal(t, expected.notify, notifyRes)
	assert.Equal(t, []uint64{1, 2}, userIDs)
}

func TestMessageInteractor_Success_Update(t *testing.T) {
	ctx := context.Background()

	repo := &mockMessageRepository{}
	userRepo := &mockUserRepository{}
	notifyRepo := &mockNotifyRepository{}
	msgPre := &mockMessagePresenter{}
	notifyPre := &mockNotifyPresenter{}
	interactor := usecase.NewMessageInteractor(repo, userRepo, notifyRepo, msgPre, notifyPre)

	err := interactor.Update(ctx, "1", "Hello, World!")

	assert.NoError(t, err)
}

func TestMessageInteractor_Success_Delete(t *testing.T) {
	ctx := context.Background()

	repo := &mockMessageRepository{}
	userRepo := &mockUserRepository{}
	notifyRepo := &mockNotifyRepository{}
	msgPre := &mockMessagePresenter{}
	notifyPre := &mockNotifyPresenter{}
	interactor := usecase.NewMessageInteractor(repo, userRepo, notifyRepo, msgPre, notifyPre)

	err := interactor.Delete(ctx, "1")

	assert.NoError(t, err)
}

func TestMessageInteractor_Failed_GetMessagesByChannelID(t *testing.T) {
	ctx := context.WithValue(context.Background(), GetMessagesByChannelIDFailedValue, true)

	repo := &mockMessageRepository{}
	userRepo := &mockUserRepository{}
	notifyRepo := &mockNotifyRepository{}
	msgPre := &mockMessagePresenter{}
	notifyPre := &mockNotifyPresenter{}
	interactor := usecase.NewMessageInteractor(repo, userRepo, notifyRepo, msgPre, notifyPre)

	res, err := interactor.GetMessagesByChannelID(ctx, 1)

	assert.Error(t, err)
	assert.Nil(t, res)
}

func TestMessageInteractor_Failed_Create__Create_Message_Failed(t *testing.T) {
	ctx := context.WithValue(context.Background(), CreateMessageFailedValue, true)

	repo := &mockMessageRepository{}
	userRepo := &mockUserRepository{}
	notifyRepo := &mockNotifyRepository{}
	msgPre := &mockMessagePresenter{}
	notifyPre := &mockNotifyPresenter{}
	interactor := usecase.NewMessageInteractor(repo, userRepo, notifyRepo, msgPre, notifyPre)

	res, notifyRes, userIDs, err := interactor.Create(ctx, 1, 1, "Hello, World!")

	assert.Error(t, err)
	assert.Nil(t, res)
	assert.Nil(t, notifyRes)
	assert.Nil(t, userIDs)
	assert.Equal(t, "failed to create message", err.Error())
}

func TestMessageInteractor_Failed_CreateReply__Target_Message_Not_Found(t *testing.T) {
	ctx := context.WithValue(context.Background(), FindByIDFailedValue, true)

	repo := &mockMessageRepository{}
	userRepo := &mockUserRepository{}
	notifyRepo := &mockNotifyRepository{}
	msgPre := &mockMessagePresenter{}
	notifyPre := &mockNotifyPresenter{}
	interactor := usecase.NewMessageInteractor(repo, userRepo, notifyRepo, msgPre, notifyPre)

	res, notifyRes, userIDs, err := interactor.CreateReply(ctx, 1, 1, "1", "Hello, World!")

	assert.Error(t, err)
	assert.Nil(t, res)
	assert.Nil(t, notifyRes)
	assert.Nil(t, userIDs)
	assert.Equal(t, "failed to get message by id", err.Error())
}

func TestMessageInteractor_Failed_CreateReply__Create_Reply_Failed(t *testing.T) {
	ctx := context.WithValue(context.Background(), CreateReplyFailedValue, true)

	repo := &mockMessageRepository{}
	userRepo := &mockUserRepository{}
	notifyRepo := &mockNotifyRepository{}
	msgPre := &mockMessagePresenter{}
	notifyPre := &mockNotifyPresenter{}
	interactor := usecase.NewMessageInteractor(repo, userRepo, notifyRepo, msgPre, notifyPre)

	res, notifyRes, userIDs, err := interactor.CreateReply(ctx, 1, 1, "1", "Hello, World!")

	assert.Error(t, err)
	assert.Nil(t, res)
	assert.Nil(t, notifyRes)
	assert.Nil(t, userIDs)
	assert.Equal(t, "failed to create reply", err.Error())
}

func TestMessageInteractor_Failed_CreateReply__Get_Message_Failed(t *testing.T) {
	ctx := context.WithValue(context.Background(), FindByIDFailedValue, true)

	repo := &mockMessageRepository{}
	userRepo := &mockUserRepository{}
	notifyRepo := &mockNotifyRepository{}
	msgPre := &mockMessagePresenter{}
	notifyPre := &mockNotifyPresenter{}
	interactor := usecase.NewMessageInteractor(repo, userRepo, notifyRepo, msgPre, notifyPre)

	res, notifyRes, userIDs, err := interactor.CreateReply(ctx, 1, 1, "1", "Hello, World!")

	assert.Error(t, err)
	assert.Nil(t, res)
	assert.Nil(t, notifyRes)
	assert.Nil(t, userIDs)
	assert.Equal(t, "failed to get message by id", err.Error())
}

func TestMessageInteractor_Failed_Update__Target_Message_Not_Found(t *testing.T) {
	ctx := context.WithValue(context.Background(), FindByIDFailedValue, true)

	repo := &mockMessageRepository{}
	userRepo := &mockUserRepository{}
	notifyRepo := &mockNotifyRepository{}
	msgPre := &mockMessagePresenter{}
	notifyPre := &mockNotifyPresenter{}
	interactor := usecase.NewMessageInteractor(repo, userRepo, notifyRepo, msgPre, notifyPre)

	err := interactor.Update(ctx, "1", "Hello, World!")

	assert.Error(t, err)
	assert.Equal(t, "failed to get message by id", err.Error())
}

func TestMessageInteractor_Failed_Update__Update_Message_Failed(t *testing.T) {
	ctx := context.WithValue(context.Background(), UpdateMessageFailedValue, true)

	repo := &mockMessageRepository{}
	userRepo := &mockUserRepository{}
	notifyRepo := &mockNotifyRepository{}
	msgPre := &mockMessagePresenter{}
	notifyPre := &mockNotifyPresenter{}
	interactor := usecase.NewMessageInteractor(repo, userRepo, notifyRepo, msgPre, notifyPre)

	err := interactor.Update(ctx, "1", "Hello, World!")

	assert.Error(t, err)
	assert.Equal(t, "failed to update message", err.Error())
}

func TestMessageInteractor_Failed_Delete(t *testing.T) {
	ctx := context.WithValue(context.Background(), DeleteMessageFailedValue, true)

	repo := &mockMessageRepository{}
	userRepo := &mockUserRepository{}
	notifyRepo := &mockNotifyRepository{}
	msgPre := &mockMessagePresenter{}
	notifyPre := &mockNotifyPresenter{}
	interactor := usecase.NewMessageInteractor(repo, userRepo, notifyRepo, msgPre, notifyPre)

	err := interactor.Delete(ctx, "1")

	assert.Error(t, err)
}

type mockMessageRepository struct{}

func (m *mockMessageRepository) FindByChannelID(ctx context.Context, channelID uint64) ([]*model.Message, error) {
	if ctx.Value(GetMessagesByChannelIDFailedValue) != nil {
		return nil, fmt.Errorf("failed to get messages by channel id")
	}

	return []*model.Message{
		{
			ID: "1",
			User: &model.User{
				ID:       1,
				Name:     "User",
				ImageURL: "https://example.com/image.png",
			},
			Content:        "Hello, World!",
			ReplyToMessage: &model.Message{},
			CreatedAt:      time.Time{},
			UpdatedAt:      time.Time{},
		},
	}, nil
}

func (m *mockMessageRepository) FindByID(ctx context.Context, id string) (*model.Message, error) {
	if ctx.Value(FindByIDFailedValue) != nil {
		return nil, fmt.Errorf("failed to get message by id")
	}

	return &model.Message{
		ID: "1",
		User: &model.User{
			ID:       1,
			Name:     "User",
			ImageURL: "https://example.com/image.png",
		},
		ChannelID:      1,
		ReplyToMessage: &model.Message{},
		Content:        "Hello, World!",
		CreatedAt:      time.Time{},
		UpdatedAt:      time.Time{},
	}, nil
}

func (m *mockMessageRepository) Create(ctx context.Context, message *model.Message) error {
	if ctx.Value(CreateMessageFailedValue) != nil {
		return fmt.Errorf("failed to create message")
	}

	return nil
}

func (m *mockMessageRepository) CreateReply(ctx context.Context, message *model.Message) error {
	if ctx.Value(CreateReplyFailedValue) != nil {
		return fmt.Errorf("failed to create reply")
	}

	return nil
}

func (m *mockMessageRepository) Update(ctx context.Context, message *model.Message) error {
	if ctx.Value(UpdateMessageFailedValue) != nil {
		return fmt.Errorf("failed to update message")
	}

	return nil
}

func (m *mockMessageRepository) Delete(ctx context.Context, id string) error {
	if ctx.Value(DeleteMessageFailedValue) != nil {
		return fmt.Errorf("failed to delete message")
	}

	return nil
}

type mockNotifyRepository struct{}

func (m *mockNotifyRepository) BulkCreate(ctx context.Context, notify []model.Notify) error {
	return nil
}

func (m *mockNotifyRepository) FindByUserID(ctx context.Context, userID uint64) ([]*model.Notify, error) {
	return nil, nil
}

type mockMessagePresenter struct{}

func (m *mockMessagePresenter) GenerateGetMessagesByChannelIDResponse(messages []*entity.Message) *presenter.GetMessagesByChannelIDResponse {
	messagesResponse := &presenter.GetMessagesByChannelIDResponse{
		Messages: []*presenter.Message{},
	}
	for _, message := range messages {
		var replyToMessage *presenter.ReplyToMessage = nil
		messagesResponse.Messages = append(messagesResponse.Messages, &presenter.Message{
			ID: message.ID,
			User: &presenter.User{
				ID:       strconv.FormatUint(message.User.ID, 10),
				Name:     message.User.Name,
				ImageURL: message.User.ImageURL,
			},
			Content:        message.Content,
			ReplyToMessage: replyToMessage,
			CreatedAt:      message.CreatedAt.String(),
			UpdatedAt:      message.UpdatedAt.String(),
		})
	}

	return messagesResponse
}

func (m *mockMessagePresenter) GenerateCreateMessageResponse(message *entity.Message) *presenter.CreateMessageResponse {
	var replyToMessage *presenter.ReplyToMessage = nil

	return &presenter.CreateMessageResponse{
		Message: &presenter.Message{
			ID: message.ID,
			User: &presenter.User{
				ID:       strconv.FormatUint(message.User.ID, 10),
				Name:     message.User.Name,
				ImageURL: message.User.ImageURL,
			},
			Content:        message.Content,
			ReplyToMessage: replyToMessage,
			CreatedAt:      message.CreatedAt.String(),
			UpdatedAt:      message.UpdatedAt.String(),
		},
		ChannelID: strconv.FormatUint(message.ChannelID, 10),
	}
}

type mockNotifyPresenter struct{}

func (m *mockNotifyPresenter) GenerateNotifyMessageResponse(message *entity.Message) *presenter.NotifyMessageResponse {
	return &presenter.NotifyMessageResponse{
		NotifyMeta: presenter.NotifyMeta{
			TypeID: config.NOTIFY_MESSAGE,
		},
		Payload: presenter.Payload{
			Message: &presenter.Message{
				ID: message.ID,
				User: &presenter.User{
					ID:       strconv.FormatUint(message.User.ID, 10),
					Name:     message.User.Name,
					ImageURL: message.User.ImageURL,
				},
				Content:        message.Content,
				ReplyToMessage: nil,
				CreatedAt:      message.CreatedAt.String(),
				UpdatedAt:      message.UpdatedAt.String(),
			},
			ChannelID: strconv.FormatUint(message.ChannelID, 10),
		},
	}
}
