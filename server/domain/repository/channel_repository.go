package repository

import (
	"context"

	"github.com/saitamau-maximum/meline/models"
)

type IChannelRepository interface {
	FindByID(ctx context.Context, id uint64) (*model.Channel, error)
	Create(ctx context.Context, channel *model.Channel) (uint64, error)
	Update(ctx context.Context, channel *model.Channel) error
	Delete(ctx context.Context, id uint64) error
}
