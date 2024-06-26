package mysql

import (
	"context"
	"sync"

	"github.com/uptrace/bun"

	"github.com/saitamau-maximum/meline/domain/repository"
	model "github.com/saitamau-maximum/meline/models"
)

type UserRepository struct {
	db *bun.DB
	mu sync.RWMutex
}

func NewUserRepository(db *bun.DB) repository.IUserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) FindByID(ctx context.Context, id uint64) (*model.User, error) {
	var user model.User

	if err := r.db.NewSelect().Model(&model.User{}).Where("id = ?", id).Where("deleted_at IS NULL").Scan(ctx, &user); err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) FindByProviderID(ctx context.Context, providerID string) (*model.User, error) {
	var user model.User

	if err := r.db.NewSelect().Model(&model.User{}).Where("provider_id = ?", providerID).Where("deleted_at IS NULL").Scan(ctx, &user); err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) Create(ctx context.Context, user *model.User) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, err := r.db.NewInsert().Model(user).Exec(ctx); err != nil {
		return err
	}

	return nil
}

func (r *UserRepository) FindChannelsByUserID(ctx context.Context, userID uint64) ([]*model.Channel, error) {
	var user = &model.User{}

	// UsersリレーションにuserIdが含まれるchannelを取得する
	if err := r.db.NewSelect().Model(user).Where("id = ?", userID).Relation("Channels").Scan(ctx); err != nil {
		return nil, err
	}

	return user.Channels, nil
}

func (r *UserRepository) IsUserExists(ctx context.Context, userID uint64) (bool, error) {
	isExist, err := r.db.NewSelect().Model(&model.User{}).Where("id = ?", userID).Exists(ctx)
	if err != nil {
		return false, err
	}

	return isExist, nil
}

func (r *UserRepository) FindByChannelID(ctx context.Context, channelID uint64) ([]*model.User, error) {
	var users []*model.User

	if err := r.db.NewSelect().Model(&users).Column("user.*").Join("JOIN channel_users ON channel_users.user_id = user.id").Where("channel_users.channel_id = ?", channelID).Scan(ctx); err != nil {
		return nil, err
	}

	return users, nil
}
