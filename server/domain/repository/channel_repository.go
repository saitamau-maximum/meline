package repository

import (
	"context"

	"github.com/saitamau-maximum/meline/models"
)

type IChannelRepository interface {
	FindAll(ctx context.Context) ([]*model.Channel, error)
	FindByID(ctx context.Context, id uint64) (*model.Channel, error)
	FindByUserID(ctx context.Context, userID uint64) ([]*model.Channel, error)
	Create(ctx context.Context, channel *model.Channel) error
	Update(ctx context.Context, channel *model.Channel) error
	Delete(ctx context.Context, id uint64) error
}
