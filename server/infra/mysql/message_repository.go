package mysql

import (
	"context"
	"sync"

	"github.com/saitamau-maximum/meline/domain/repository"
	model "github.com/saitamau-maximum/meline/models"
	"github.com/uptrace/bun"
)

type messageRepository struct {
	db *bun.DB
	mu sync.RWMutex
}

func NewMessageRepository(db *bun.DB) repository.IMessageRepository {
	return &messageRepository{db: db}
}

func (r *messageRepository) FindByChannelID(ctx context.Context, channelID uint64) ([]*model.Message, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	var messages []*model.Message
	err := r.db.NewSelect().Model(&messages).Where("message.channel_id = ?", channelID).Relation("ReplyToMessage").Relation("ReplyToMessage.User").Relation("User").Order("created_at ASC").Scan(ctx)
	if err != nil {
		return nil, err
	}

	return messages, nil
}

func (r *messageRepository) FindByID(ctx context.Context, id string) (*model.Message, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	message := &model.Message{}
	err := r.db.NewSelect().Model(message).Where("message.id = ?", id).Relation("ReplyToMessage").Relation("ReplyToMessage.User").Scan(ctx)
	if err != nil {
		return nil, err
	}

	return message, nil
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

func (r *messageRepository) CreateReply(ctx context.Context, message *model.Message) error {
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

	_, err := r.db.NewUpdate().Model(message).Where("message.id = ?", message.ID).Exec(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (r *messageRepository) Delete(ctx context.Context, id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	_, err := r.db.NewDelete().Model(&model.Message{}).Where("message.id = ?", id).Exec(ctx)
	if err != nil {
		return err
	}

	return nil
}
