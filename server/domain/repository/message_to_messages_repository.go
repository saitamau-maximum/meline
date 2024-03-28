package repository

import (
	"context"

	"github.com/saitamau-maximum/meline/models"
)

type IMessageToMessagesRepository interface {
	Create(ctx context.Context, messageToMessages *model.MessageToMessages) error
	DeleteByMessageID(ctx context.Context, MessageID string) error
}
