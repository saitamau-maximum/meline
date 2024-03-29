package usecase

import (
	"context"

	"github.com/saitamau-maximum/meline/domain/entity"
	"github.com/saitamau-maximum/meline/domain/repository"
	"github.com/saitamau-maximum/meline/models"
	"github.com/saitamau-maximum/meline/usecase/presenter"
)

type IMessageInteractor interface {
	GetMessagesByChannelID(ctx context.Context, channelID uint64) (*presenter.GetMessagesByChannelIDResponse, error)
	Create(ctx context.Context, userID, channelID uint64, content string) (*entity.Message, error)
	CreateReply(ctx context.Context, userID, channelID uint64, parentMessageID string, content string) (*entity.Message, error)
	Update(ctx context.Context, id string, content string) error
	Delete(ctx context.Context, id string) error
}

type messageInteractor struct {
	messageRepository           repository.IMessageRepository
	messageToMessagesRepository repository.IMessageToMessagesRepository
	messagePresenter            presenter.IMessagePresenter
}

func NewMessageInteractor(messageRepository repository.IMessageRepository, messageToMessagesRepository repository.IMessageToMessagesRepository, messagePresenter presenter.IMessagePresenter) IMessageInteractor {
	return &messageInteractor{
		messageRepository:           messageRepository,
		messageToMessagesRepository: messageToMessagesRepository,
		messagePresenter:            messagePresenter,
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

func (i *messageInteractor) Create(ctx context.Context, userID, channelID uint64, content string) (*entity.Message, error) {
	message := model.NewMessageModel(channelID, userID, content)

	if err := i.messageRepository.Create(ctx, message); err != nil {
		return nil, err
	}

	createdMsg, err := i.messageRepository.FindByID(ctx, message.ID)
	if err != nil {
		return nil, err
	}

	return createdMsg.ToMessageEntity(), nil
}

func (i *messageInteractor) CreateReply(ctx context.Context, userID, channelID uint64, parentMessageID string, content string) (*entity.Message, error) {
	message := model.NewMessageModel(channelID, userID, content)

	if err := i.messageRepository.Create(ctx, message); err != nil {
		return nil, err
	}

	messageToMessages := model.NewMessageToMessagesModel(message.ID, parentMessageID)

	if err := i.messageToMessagesRepository.Create(ctx, messageToMessages); err != nil {
		return nil, err
	}

	createdMsg, err := i.messageRepository.FindByID(ctx, message.ID)
	if err != nil {
		return nil, err
	}

	return createdMsg.ToMessageEntity(), nil
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

	if err := i.messageToMessagesRepository.DeleteByMessageID(ctx, id); err != nil {
		return err
	}

	return nil
}
