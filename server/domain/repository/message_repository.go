package repository

import (
	"context"
	
	"github.com/saitamau-maximum/meline/models"
)

type MessageRepository interface {
	FindByID(ctx context.Context, id uint64) (*model.Message, error)
	FindByChannelID(ctx context.Context, channelID uint64) ([]*model.Message, error)
	Create(ctx context.Context, message *model.Message) error
	Update(ctx context.Context, message *model.Message) error
	Delete(ctx context.Context, id uint64) error
}
