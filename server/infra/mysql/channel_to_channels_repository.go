package mysql

import (
	"context"
	"sync"

	"github.com/saitamau-maximum/meline/domain/repository"
	"github.com/saitamau-maximum/meline/models"
	"github.com/uptrace/bun"
)

type ChannelToChannelsRepository struct {
	db *bun.DB
	mu sync.RWMutex
}

func NewChannelToChannelsRepository(db *bun.DB) repository.IChannelToChannelsRepository {
	return &ChannelToChannelsRepository{
		db: db,
	}
}

func (r *ChannelToChannelsRepository) Create(ctx context.Context, channel *model.ChannelToChannels) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	_, err := r.db.NewInsert().Model(channel).Exec(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (r *ChannelToChannelsRepository) DeleteFromParentChannelID(ctx context.Context, parentChannelID uint64) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	_, err := r.db.NewDelete().Model(&model.ChannelToChannels{}).Where("parent_channel_id = ?", parentChannelID).Exec(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (r *ChannelToChannelsRepository) DeleteFromChildChannelID(ctx context.Context, childChannelID uint64) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	_, err := r.db.NewDelete().Model(&model.ChannelToChannels{}).Where("child_channel_id = ?", childChannelID).Exec(ctx)
	if err != nil {
		return err
	}

	return nil
}
