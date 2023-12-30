package usecase

import (
	"context"

	"github.com/saitamau-maximum/meline/domain/repository"
	"github.com/saitamau-maximum/meline/models"
)

type IMessageInteractor interface {
	Create(ctx context.Context, userID, channelID uint64, replyToID, threadID, content string) error
	Update(ctx context.Context, id string, content string) error
	Delete(ctx context.Context, id string) error
}

type messageInteractor struct {
	messageRepository repository.IMessageRepository
}

func NewMessageInteractor(messageRepository repository.IMessageRepository) IMessageInteractor {
	return &messageInteractor{
		messageRepository: messageRepository,
	}
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
