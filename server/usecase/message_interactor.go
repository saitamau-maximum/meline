package usecase

import (
	"context"

	"github.com/saitamau-maximum/meline/config"
	"github.com/saitamau-maximum/meline/domain/entity"
	"github.com/saitamau-maximum/meline/domain/repository"
	model "github.com/saitamau-maximum/meline/models"
	"github.com/saitamau-maximum/meline/usecase/presenter"
)

type IMessageInteractor interface {
	GetMessagesByChannelID(ctx context.Context, channelID uint64) (*presenter.GetMessagesByChannelIDResponse, error)
	Create(ctx context.Context, userID, channelID uint64, content string) (*presenter.CreateMessageResponse, *presenter.NotifyMessageResponse, []uint64, error)
	CreateReply(ctx context.Context, userID, channelID uint64, parentMessageID string, content string) (*presenter.CreateMessageResponse, *presenter.NotifyMessageResponse, []uint64, error)
	Update(ctx context.Context, id string, content string) error
	Delete(ctx context.Context, id string) error
}

type messageInteractor struct {
	messageRepository repository.IMessageRepository
	userRepository    repository.IUserRepository
	notitfyRepository repository.INotifyRepository
	messagePresenter  presenter.IMessagePresenter
	notifyPresenter   presenter.INotifyPresenter
}

func NewMessageInteractor(messageRepository repository.IMessageRepository, userRepository repository.IUserRepository, notifyRepository repository.INotifyRepository, messagePresenter presenter.IMessagePresenter, notifPresenter presenter.INotifyPresenter) IMessageInteractor {
	return &messageInteractor{
		messageRepository: messageRepository,
		userRepository:    userRepository,
		notitfyRepository: notifyRepository,
		messagePresenter:  messagePresenter,
		notifyPresenter:   notifPresenter,
	}
}

func (i *messageInteractor) GetMessagesByChannelID(ctx context.Context, channelID uint64) (*presenter.GetMessagesByChannelIDResponse, error) {
	messages, err := i.messageRepository.FindByChannelID(ctx, channelID)
	if err != nil {
		return nil, err
	}

	entitiedMessages := make([]*entity.Message, 0)
	for _, message := range messages {
		entitiedMessages = append(entitiedMessages, message.ToMessageEntity())
	}

	return i.messagePresenter.GenerateGetMessagesByChannelIDResponse(entitiedMessages), nil
}

func (i *messageInteractor) Create(ctx context.Context, userID, channelID uint64, content string) (*presenter.CreateMessageResponse, *presenter.NotifyMessageResponse, []uint64, error) {
	message := model.NewMessageModel(channelID, userID, content)

	if err := i.messageRepository.Create(ctx, message); err != nil {
		return nil, nil, nil, err
	}

	createdMsg, err := i.messageRepository.FindByID(ctx, message.ID)
	if err != nil {
		return nil, nil, nil, err
	}

	users, err := i.userRepository.FindByChannelID(ctx, channelID)
	if err != nil {
		return nil, nil, nil, err
	}

	userIDs := make([]uint64, 0)
	for _, user := range users {
		userIDs = append(userIDs, user.ID)
	}

	notifies := make([]model.Notify, 0)
	for _, user := range users {
		notify := model.Notify{
			UserID:    user.ID,
			TypeID:    config.NOTIFY_MESSAGE,
			MessageID: createdMsg.ID,
		}

		notifies = append(notifies, notify)
	}

	if err := i.notitfyRepository.BulkCreate(ctx, notifies); err != nil {
		return nil, nil, nil, err
	}

	return i.messagePresenter.GenerateCreateMessageResponse(createdMsg.ToMessageEntity()), i.notifyPresenter.GenerateNotifyMessageResponse(createdMsg.ToMessageEntity()), userIDs, nil
}

func (i *messageInteractor) CreateReply(ctx context.Context, userID, channelID uint64, parentMessageID string, content string) (*presenter.CreateMessageResponse, *presenter.NotifyMessageResponse, []uint64, error) {
	message := model.NewMessageModel(channelID, userID, content)
	message.ReplyToMessageID = parentMessageID

	if err := i.messageRepository.CreateReply(ctx, message); err != nil {
		return nil, nil, nil, err
	}

	createdMsg, err := i.messageRepository.FindByID(ctx, message.ID)
	if err != nil {
		return nil, nil, nil, err
	}

	users, err := i.userRepository.FindByChannelID(ctx, channelID)
	if err != nil {
		return nil, nil, nil, err
	}

	userIDs := make([]uint64, 0)
	for _, user := range users {
		userIDs = append(userIDs, user.ID)
	}

	notifies := make([]model.Notify, 0)
	for _, user := range users {
		notify := model.Notify{
			UserID:    user.ID,
			TypeID:    config.NOTIFY_MESSAGE,
			MessageID: createdMsg.ID,
		}

		notifies = append(notifies, notify)
	}

	if err := i.notitfyRepository.BulkCreate(ctx, notifies); err != nil {
		return nil, nil, nil, err
	}

	return i.messagePresenter.GenerateCreateMessageResponse(createdMsg.ToMessageEntity()), i.notifyPresenter.GenerateNotifyMessageResponse(createdMsg.ToMessageEntity()), userIDs, nil
}

func (i *messageInteractor) Update(ctx context.Context, id string, content string) error {
	message, err := i.messageRepository.FindByID(ctx, id)
	if err != nil {
		return err
	}

	message.Content = content

	if err := i.messageRepository.Update(ctx, message); err != nil {
		return err
	}

	return nil
}

func (i *messageInteractor) Delete(ctx context.Context, id string) error {
	if err := i.messageRepository.Delete(ctx, id); err != nil {
		return err
	}

	return nil
}
