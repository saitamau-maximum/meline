package usecase

import (
	"context"

	"github.com/saitamau-maximum/meline/domain/entity"
	"github.com/saitamau-maximum/meline/domain/repository"
	"github.com/saitamau-maximum/meline/models"
	"github.com/saitamau-maximum/meline/usecase/presenter"
)

type IMessageInteractor interface {
	FindByID(ctx context.Context, id string) (*presenter.GetMessageByIDResponse, error)
	Create(ctx context.Context, userID, channelID uint64, replyToID, threadID, content string) error
	Update(ctx context.Context, id string, content string) error
	Delete(ctx context.Context, id string) error
}

type messageInteractor struct {
	messageRepository repository.IMessageRepository
	messagePresenter  presenter.IMessagePresenter
}

func NewMessageInteractor(messageRepository repository.IMessageRepository, messagePresenter presenter.IMessagePresenter) IMessageInteractor {
	return &messageInteractor{
		messageRepository: messageRepository,
		messagePresenter:  messagePresenter,
	}
}

func (i *messageInteractor) FindByID(ctx context.Context, id string) (*presenter.GetMessageByIDResponse, error) {
	message, err := i.messageRepository.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return i.messagePresenter.GenerateGetMessageByIDResponse(message.ToMessageEntity()), nil
}

func (i *messageInteractor) Create(ctx context.Context, userID, channelID uint64, replyToID, threadID, content string) error {
	message := model.NewMessageModel(userID, channelID, replyToID, threadID, content)

	if err := i.messageRepository.Create(ctx, message); err != nil {
		return err
	}

	return nil
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
