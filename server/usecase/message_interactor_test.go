package usecase_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/saitamau-maximum/meline/domain/entity"
	"github.com/saitamau-maximum/meline/models"
	"github.com/saitamau-maximum/meline/usecase"
	"github.com/saitamau-maximum/meline/usecase/presenter"

	"github.com/stretchr/testify/assert"
)

type GetMessagesByChannelIDFailed string
type CreateMessageFailed string
type CreateReplyFailed string
type UpdateMessageFailed string
type DeleteMessageFailed string

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
	repoMessageToMessages := &mockMessageToMessagesRepository{}
	pre := &mockMessagePresenter{}
	interactor := usecase.NewMessageInteractor(repo, repoMessageToMessages, pre)

	res, err := interactor.GetMessagesByChannelID(ctx, 1)

	expected := &presenter.GetMessagesByChannelIDResponse{
		Messages: []*presenter.Message{
			{
				ID:      "1",
				User:    &presenter.User{
					ID:       1,
					Name:     "User",
					ImageURL: "https://example.com/image.png",
				},
				Content: "Hello, World!",
				ReplyToMessage: []*presenter.ReplyToMessage{},
				CreatedAt: "0001-01-01 00:00:00 +0000 UTC",
				UpdatedAt: "0001-01-01 00:00:00 +0000 UTC",
			},
		},
	}

	assert.NoError(t, err)
	assert.Equal(t, expected, res)
}

func TestMessageInteractor_Success_Create(t *testing.T) {
	ctx := context.Background()

	repo := &mockMessageRepository{}
	repoMessageToMessages := &mockMessageToMessagesRepository{}
	pre := &mockMessagePresenter{}
	interactor := usecase.NewMessageInteractor(repo, repoMessageToMessages, pre)

	res, err := interactor.Create(ctx, 1, 1, "Hello, World!")

	expected := &entity.Message{
		ID:             "1",
		Channel:        nil,
		User:           &entity.User{
			ID:       1,
			Name:     "User",
			ImageURL: "https://example.com/image.png",
		},
		ReplyToMessage: []*entity.Message{},
		Content:        "Hello, World!",
		CreatedAt:      time.Time{},
		UpdatedAt:      time.Time{},
	}

	assert.NoError(t, err)
	assert.Equal(t, expected, res)
}

func TestMessageInteractor_Success_CreateReply(t *testing.T) {
	ctx := context.Background()

	repo := &mockMessageRepository{}
	repoMessageToMessages := &mockMessageToMessagesRepository{}
	pre := &mockMessagePresenter{}
	interactor := usecase.NewMessageInteractor(repo, repoMessageToMessages, pre)

	res, err := interactor.CreateReply(ctx, 1, 1, "1", "Hello, World!")

	expected := &entity.Message{
		ID:             "1",
		Channel:        nil,
		User:           &entity.User{
			ID:       1,
			Name:     "User",
			ImageURL: "https://example.com/image.png",
		},
		ReplyToMessage: []*entity.Message{},
		Content:        "Hello, World!",
		CreatedAt:      time.Time{},
		UpdatedAt:      time.Time{},
	}

	assert.NoError(t, err)
	assert.Equal(t, expected, res)
}

func TestMessageInteractor_Success_Update(t *testing.T) {
	ctx := context.Background()

	repo := &mockMessageRepository{}
	repoMessageToMessages := &mockMessageToMessagesRepository{}
	pre := &mockMessagePresenter{}
	interactor := usecase.NewMessageInteractor(repo, repoMessageToMessages, pre)

	err := interactor.Update(ctx, "1", "Hello, World!")

	assert.NoError(t, err)
}

func TestMessageInteractor_Success_Delete(t *testing.T) {
	ctx := context.Background()

	repo := &mockMessageRepository{}
	repoMessageToMessages := &mockMessageToMessagesRepository{}
	pre := &mockMessagePresenter{}
	interactor := usecase.NewMessageInteractor(repo, repoMessageToMessages, pre)

	err := interactor.Delete(ctx, "1")
	
	assert.NoError(t, err)
}

func TestMessageInteractor_Failed_GetMessagesByChannelID(t *testing.T) {
	ctx := context.WithValue(context.Background(), GetMessagesByChannelIDFailedValue, true)

	repo := &mockMessageRepository{}
	repoMessageToMessages := &mockMessageToMessagesRepository{}
	pre := &mockMessagePresenter{}
	interactor := usecase.NewMessageInteractor(repo, repoMessageToMessages, pre)

	res, err := interactor.GetMessagesByChannelID(ctx, 1)

	assert.Error(t, err)
	assert.Nil(t, res)
}

func TestMessageInteractor_Failed_Create(t *testing.T) {
	ctx := context.WithValue(context.Background(), CreateMessageFailedValue, true)
	
	repo := &mockMessageRepository{}
	repoMessageToMessages := &mockMessageToMessagesRepository{}
	pre := &mockMessagePresenter{}
	interactor := usecase.NewMessageInteractor(repo, repoMessageToMessages, pre)

	res, err := interactor.Create(ctx, 1, 1, "Hello, World!")

	assert.Error(t, err)
	assert.Nil(t, res)
}

func TestMessageInteractor_Failed_CreateReply(t *testing.T) {
	ctx := context.WithValue(context.Background(), CreateReplyFailedValue, true)
	
	repo := &mockMessageRepository{}
	repoMessageToMessages := &mockMessageToMessagesRepository{}
	pre := &mockMessagePresenter{}
	interactor := usecase.NewMessageInteractor(repo, repoMessageToMessages, pre)

	res, err := interactor.CreateReply(ctx, 1, 1, "1", "Hello, World!")

	assert.Error(t, err)
	assert.Nil(t, res)
}

func TestMessageInteractor_Failed_Update(t *testing.T) {
	ctx := context.WithValue(context.Background(), UpdateMessageFailedValue, true)

	repo := &mockMessageRepository{}
	repoMessageToMessages := &mockMessageToMessagesRepository{}
	pre := &mockMessagePresenter{}
	interactor := usecase.NewMessageInteractor(repo, repoMessageToMessages, pre)

	err := interactor.Update(ctx, "1", "Hello, World!")

	assert.Error(t, err)
}

func TestMessageInteractor_Failed_Delete(t *testing.T) {
	ctx := context.WithValue(context.Background(), DeleteMessageFailedValue, true)

	repo := &mockMessageRepository{}
	repoMessageToMessages := &mockMessageToMessagesRepository{}
	pre := &mockMessagePresenter{}
	interactor := usecase.NewMessageInteractor(repo, repoMessageToMessages, pre)

	err := interactor.Delete(ctx, "1")

	assert.Error(t, err)
}

type mockMessageRepository struct {}

func (m *mockMessageRepository) FindByChannelID(ctx context.Context, channelID uint64) ([]*model.Message, error) {
	if ctx.Value(GetMessagesByChannelIDFailedValue) != nil {
		return nil, fmt.Errorf("failed to get messages by channel id")
	}

	return []*model.Message{
		{
			ID: "1",
			User: &model.User{
				ID: 1,
				Name: "User",
				ImageURL: "https://example.com/image.png",
			},
			Content: "Hello, World!",
			ReplyToMessage: []*model.Message{},
			CreatedAt: time.Time{},
			UpdatedAt: time.Time{},
		},
	}, nil
}

func (m *mockMessageRepository) FindByID(ctx context.Context, id string) (*model.Message, error) {
	if ctx.Value(CreateFailedValue) != nil {
		return nil, fmt.Errorf("failed to get message by id")
	}
	
	return &model.Message{
		ID: "1",
		User: &model.User{
			ID: 1,
			Name: "User",
			ImageURL: "https://example.com/image.png",
		},
		ReplyToMessage: []*model.Message{},
		Content: "Hello, World!",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	}, nil
}

func (m *mockMessageRepository) Create(ctx context.Context, message *model.Message) error {
	if ctx.Value(CreateMessageFailedValue) != nil {
		return fmt.Errorf("failed to create message")
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

type mockMessageToMessagesRepository struct {}

func (m *mockMessageToMessagesRepository) Create(ctx context.Context, messageToMessages *model.MessageToMessages) error {
	if ctx.Value(CreateReplyFailedValue) != nil {
		return fmt.Errorf("failed to create reply")
	}

	return nil
}

func (m *mockMessageToMessagesRepository) DeleteByMessageID(ctx context.Context, messageID string) error {
	if ctx.Value(DeleteMessageFailedValue) != nil {
		return fmt.Errorf("failed to delete message")
	}

	return nil
}

type mockMessagePresenter struct {}

func (m *mockMessagePresenter) GenerateGetMessagesByChannelIDResponse(messages []*entity.Message) *presenter.GetMessagesByChannelIDResponse {
	messagesResponse := &presenter.GetMessagesByChannelIDResponse{
		Messages: []*presenter.Message{},
	}
	for _, message := range messages {
		replyToMessages := make([]*presenter.ReplyToMessage, 0)
		for _, replyToMessage := range message.ReplyToMessage {
			replyToMessages = append(replyToMessages, &presenter.ReplyToMessage{
				ID:      replyToMessage.ID,
				User:    &presenter.User{
					ID:       replyToMessage.User.ID,
					Name:     replyToMessage.User.Name,
					ImageURL: replyToMessage.User.ImageURL,
				},
				Content: replyToMessage.Content,
			})
		}
		messagesResponse.Messages = append(messagesResponse.Messages, &presenter.Message{
			ID:      message.ID,
			User:    &presenter.User{
				ID:       message.User.ID,
				Name:     message.User.Name,
				ImageURL: message.User.ImageURL,
			},
			Content:        message.Content,
			ReplyToMessage: replyToMessages,
			CreatedAt:      message.CreatedAt.String(),
			UpdatedAt:      message.UpdatedAt.String(),
		})
	}

	return messagesResponse
}
