package repository

import (
	"context"

	"github.com/saitamau-maximum/meline/models"
)

type IChannelToChannelsRepository interface {
	Create(ctx context.Context, channel *model.ChannelToChannels) error
	DeleteFromParentChannelID(ctx context.Context, parentChannelID uint64) error
	DeleteFromChildChannelID(ctx context.Context, childChannelID uint64) error
}
