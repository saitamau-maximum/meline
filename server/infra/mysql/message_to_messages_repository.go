package mysql

import (
	"context"
	"sync"

	"github.com/saitamau-maximum/meline/domain/repository"
	"github.com/saitamau-maximum/meline/models"
	"github.com/uptrace/bun"
)

type messageToMessagesRepository struct {
	db *bun.DB
	mu sync.RWMutex
}

func NewMessageToMessagesRepository(db *bun.DB) repository.IMessageToMessagesRepository {
	return &messageToMessagesRepository{db: db}
}

func (r *messageToMessagesRepository) Create(ctx context.Context, messageToMessages *model.MessageToMessages) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	_, err := r.db.NewInsert().Model(messageToMessages).Exec(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (r *messageToMessagesRepository) DeleteByMessageID(ctx context.Context, messageID string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	_, err := r.db.NewDelete().Model(&model.MessageToMessages{}).WhereOr("parent_message_id = ?", messageID).WhereOr("child_message_id = ?", messageID).Exec(ctx)
	if err != nil {
		return err
	}

	return nil
}
