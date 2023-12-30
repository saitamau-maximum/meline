package mysql

import (
	"context"
	"sync"

	"github.com/saitamau-maximum/meline/domain/repository"
	"github.com/saitamau-maximum/meline/models"
	"github.com/uptrace/bun"
)

type messageRepository struct {
	db *bun.DB;
	mu sync.RWMutex;
}

func NewMessageRepository(db *bun.DB) repository.MessageRepository {
	return &messageRepository{db: db}
}

func (r *messageRepository) FindByID(ctx context.Context, id uint64) (*model.Message, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	message := &model.Message{}
	err := r.db.NewSelect().Model(message).Where("id = ?", id).Relation("Replys").Scan(ctx)
	if err != nil {
		return nil, err
	}

	return message, nil
}

func (r *messageRepository) FindByChannelID(ctx context.Context, channelID uint64) ([]*model.Message, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	var messages []*model.Message
	err := r.db.NewSelect().Model(&messages).Where("channel_id = ?", channelID).Relation("Replys").Scan(ctx)
	if err != nil {
		return nil, err
	}

	return messages, nil
}

func (r *messageRepository) Create(ctx context.Context, message *model.Message) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	_, err := r.db.NewInsert().Model(message).Exec(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (r *messageRepository) Update(ctx context.Context, message *model.Message) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	_, err := r.db.NewUpdate().Model(message).Where("id = ?", message.ID).Exec(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (r *messageRepository) Delete(ctx context.Context, id uint64) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	_, err := r.db.NewDelete().Model(&model.Message{}).Where("id = ?", id).Exec(ctx)
	if err != nil {
		return err
	}

	return nil
}
