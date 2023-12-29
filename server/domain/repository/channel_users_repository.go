package repository

import (
	"context"

	"github.com/saitamau-maximum/meline/models"
)

type IChannelUsersRepository interface {
	Create(ctx context.Context, channelUser *model.ChannelUsers) error
	Delete(ctx context.Context, channelID uint64, userID uint64) error
	DeleteByChannelID(ctx context.Context, channelID uint64) error
}
