package mysql

import (
	"context"
	"sync"

	"github.com/uptrace/bun"
	"github.com/saitamau-maximum/meline/domain/repository"
	"github.com/saitamau-maximum/meline/models"
)

type ChannelUsersRepository struct {
	db *bun.DB
	mu sync.RWMutex
}

func NewChannelUsersRepository(db *bun.DB) repository.IChannelUsersRepository {
	return &ChannelUsersRepository{
		db: db,
	}
}

func (r *ChannelUsersRepository) Create(ctx context.Context, channelUser *model.ChannelUsers) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, err := r.db.NewInsert().Model(channelUser).Exec(ctx); err != nil {
		return err
	}

	return nil
}

func (r *ChannelUsersRepository) Delete(ctx context.Context, channelID uint64, userID uint64) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, err := r.db.NewDelete().Model(&model.ChannelUsers{}).Where("channel_id = ?", channelID).Where("user_id = ?", userID).Exec(ctx); err != nil {
		return err
	}

	return nil
}

func (r *ChannelUsersRepository) DeleteByChannelID(ctx context.Context, channelID uint64) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, err := r.db.NewDelete().Model(&model.ChannelUsers{}).Where("channel_id = ?", channelID).Exec(ctx); err != nil {
		return err
	}

	return nil
}
