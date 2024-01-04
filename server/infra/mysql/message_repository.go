package mysql

import (
	"context"
	"sync"
	"time"

	"github.com/saitamau-maximum/meline/domain/repository"
	"github.com/saitamau-maximum/meline/models"
	"github.com/uptrace/bun"
)

type messageRepository struct {
	db *bun.DB
	mu sync.RWMutex
}

func NewMessageRepository(db *bun.DB) repository.IMessageRepository {
	return &messageRepository{db: db}
}

func (r *messageRepository) FindByID(ctx context.Context, id string) (*model.Message, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	message := &model.Message{}
	err := r.db.NewSelect().Model(message).Where("id = ?", id).Relation("ReplyToMessage").Relation("ReplyToMessage.User").Scan(ctx)
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

func (r *messageRepository) Update(ctx context.Context, message *model.Message) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	_, err := r.db.NewUpdate().Model(message).Where("id = ?", message.ID).Exec(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (r *messageRepository) Delete(ctx context.Context, id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	_, err := r.db.NewUpdate().Model(
		&model.Message{
			DeletedAt: time.Now(),
		}).Where("id = ?", id).Exec(ctx)
	if err != nil {
		return err
	}

	return nil
}
