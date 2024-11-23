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

func (r *ChannelRepository) FindByID(ctx context.Context, id uint64) (*model.Channel, error) {
	var channel model.Channel

	if err := r.db.NewSelect().Model(&channel).Where("id = ?", id).Relation("ChildChannels").Relation("Users").Scan(ctx); err != nil {
		return nil, err
	}

	return &channel, nil
}

func (r *ChannelRepository) Create(ctx context.Context, channel *model.Channel, userId uint64) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	res, err := r.db.NewInsert().Model(channel).Exec(ctx)
	if err != nil {
		return err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return err
	}

	channelToUsers := &model.ChannelUsers{
		ChannelID: uint64(id),
		UserID:    userId,
	}

	if _, err := r.db.NewInsert().Model(channelToUsers).Exec(ctx); err != nil {
		return err
	}

	return nil
}

func (r *ChannelRepository) CreateChildChannel(ctx context.Context, channel *model.Channel, parentChannelID uint64, userId uint64) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	res, err := r.db.NewInsert().Model(channel).Exec(ctx)
	if err != nil {
		return err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return err
	}

	channelToUsers := &model.ChannelUsers{
		ChannelID: uint64(id),
		UserID:    userId,
	}

	if _, err := r.db.NewInsert().Model(channelToUsers).Exec(ctx); err != nil {
		return err
	}

	channelToChannels := &model.ChannelToChannels{
		ParentChannelID: parentChannelID,
		ChildChannelID:  uint64(id),
	}

	if _, err := r.db.NewInsert().Model(channelToChannels).Exec(ctx); err != nil {
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

	if _, err := r.db.NewDelete().Model(&model.ChannelUsers{}).Where("channel_id = ?", id).Exec(ctx); err != nil {
		return err
	}

	if _, err := r.db.NewDelete().Model(&model.ChannelToChannels{}).WhereOr("parent_channel_id = ?", id).WhereOr("child_channel_id = ?", id).Exec(ctx); err != nil {
		return err
	}

	return nil
}

func (r *ChannelRepository) JoinChannel(ctx context.Context, channelID uint64, userID uint64) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	channelToUsers := &model.ChannelUsers{
		ChannelID: channelID,
		UserID:    userID,
	}

	if _, err := r.db.NewInsert().Model(channelToUsers).Exec(ctx); err != nil {
		return err
	}

	return nil
}

func (r *ChannelRepository) LeaveChannel(ctx context.Context, channelID uint64, userID uint64) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, err := r.db.NewDelete().Model(&model.ChannelUsers{}).Where("channel_id = ?", channelID).Where("user_id = ?", userID).Exec(ctx); err != nil {
		return err
	}

	return nil
}

func (r *ChannelRepository) FetchJoinedChannelIDs(ctx context.Context, userID uint64) ([]uint64, error) {
	var channelToUsers []model.ChannelUsers

	if err := r.db.NewSelect().Model(&channelToUsers).Where("user_id = ?", userID).Scan(ctx); err != nil {
		return nil, err
	}

	var channelIDs []uint64
	for _, channelToUser := range channelToUsers {
		channelIDs = append(channelIDs, channelToUser.ChannelID)
	}

	return channelIDs, nil
}
