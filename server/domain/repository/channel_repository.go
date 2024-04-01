package repository

import (
	"context"

	"github.com/saitamau-maximum/meline/models"
)

type IChannelRepository interface {
	FindByID(ctx context.Context, id uint64) (*model.Channel, error)
	Create(ctx context.Context, channel *model.Channel, userId uint64) error
	CreateChildChannel(ctx context.Context, channel *model.Channel, parentChannelID uint64, userId uint64) error
	Update(ctx context.Context, channel *model.Channel) error
	Delete(ctx context.Context, id uint64) error
	JoinChannel(ctx context.Context, channelID uint64, userID uint64) error
	LeaveChannel(ctx context.Context, channelID uint64, userID uint64) error
}
