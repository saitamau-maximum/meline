package usecase

import (
	"context"

	"github.com/saitamau-maximum/meline/domain/entity"
	"github.com/saitamau-maximum/meline/domain/repository"
	"github.com/saitamau-maximum/meline/models"
	"github.com/saitamau-maximum/meline/usecase/presenter"
)

type IMessageInteractor interface {
	FindByID(ctx context.Context, id uint64) (*presenter.GetMessageByIDResponse, error)
	FindByChannelID(ctx context.Context, channelID uint64) (*presenter.GetMessagesByChannelIDResponse, error)
	Create(ctx context.Context, userID, channelID uint64, replyToID, content string) error
	Update(ctx context.Context, id uint64, content string) error
	Delete(ctx context.Context, id uint64) error
}

type messageInteractor struct {
	messageRepository repository.MessageRepository
	messagePresenter  presenter.IMessagePresenter
}

func NewMessageInteractor(messageRepository repository.MessageRepository, messagePresenter presenter.IMessagePresenter) IMessageInteractor {
	return &messageInteractor{
		messageRepository: messageRepository,
		messagePresenter:  messagePresenter,
	}
}

func (i *messageInteractor) FindByID(ctx context.Context, id uint64) (*presenter.GetMessageByIDResponse, error) {
	message, err := i.messageRepository.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return i.messagePresenter.GenerateGetMessageByIDResponse(message.ToMessageEntity()), nil
}

func (i *messageInteractor) FindByChannelID(ctx context.Context, channelID uint64) (*presenter.GetMessagesByChannelIDResponse, error) {
	messages, err := i.messageRepository.FindByChannelID(ctx, channelID)
	if err != nil {
		return nil, err
	}

	entitiedMessages := make([]*entity.Message, len(messages))
	for i, message := range messages {
		entitiedMessages[i] = message.ToMessageEntity()
	}

	return i.messagePresenter.GenerateGetMessagesByChannelIDResponse(entitiedMessages), nil
}

func (i *messageInteractor) Create(ctx context.Context, userID, channelID uint64, replyToID, content string) error {
	message := model.NewMessageModel(userID, channelID, replyToID, content)

	if err := i.messageRepository.Create(ctx, message); err != nil {
		return err
	}

	return nil
}

func (i *messageInteractor) Update(ctx context.Context, id uint64, content string) error {
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

func (i *messageInteractor) Delete(ctx context.Context, id uint64) error {
	if err := i.messageRepository.Delete(ctx, id); err != nil {
		return err
	}

	return nil
}
