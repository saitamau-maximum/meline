package mysql

import (
	"context"
	"sync"

	"github.com/saitamau-maximum/meline/domain/repository"
	"github.com/saitamau-maximum/meline/models"
	"github.com/uptrace/bun"
)

type ChannelRepository struct {
	db *bun.DB
	mu sync.RWMutex
}

func NewChannelRepository(db *bun.DB) repository.IChannelRepository {
	return &ChannelRepository{
		db: db,
	}
}

func (r *ChannelRepository) FindAll(ctx context.Context) ([]*model.Channel, error) {
	var channels []*model.Channel

	if err := r.db.NewSelect().Model(&channels).Relation("users").Scan(ctx); err != nil {
		return nil, err
	}

	return channels, nil
}

func (r *ChannelRepository) FindByID(ctx context.Context, id uint64) (*model.Channel, error) {
	var channel model.Channel

	if err := r.db.NewSelect().Model(&channel).Where("id = ?", id).Relation("users").Scan(ctx); err != nil {
		return nil, err
	}

	return &channel, nil
}

func (r *ChannelRepository) FindByUserID(ctx context.Context, userID uint64) ([]*model.Channel, error) {
	var channels []*model.Channel

	if err := r.db.NewSelect().Model(&channels).Relation("users", func(q *bun.SelectQuery) *bun.SelectQuery {
		return q.Where("user_id = ?", userID)
	}).Scan(ctx); err != nil {
		return nil, err
	}

	return channels, nil
}

func (r *ChannelRepository) Create(ctx context.Context, channel *model.Channel) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, err := r.db.NewInsert().Model(channel).Exec(ctx); err != nil {
		return err
	}

	return nil
}

func (r *ChannelRepository) Update(ctx context.Context, channel *model.Channel) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, err := r.db.NewUpdate().Model(channel).Exec(ctx); err != nil {
		return err
	}

	return nil
}

func (r *ChannelRepository) Delete(ctx context.Context, id uint64) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, err := r.db.NewDelete().Model(&model.Channel{}).Where("id = ?", id).Exec(ctx); err != nil {
		return err
	}

	return nil
}
