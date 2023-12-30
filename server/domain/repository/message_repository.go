package repository

import (
	"context"
	
	"github.com/saitamau-maximum/meline/models"
)

type IMessageRepository interface {
	FindByID(ctx context.Context, id string) (*model.Message, error)
	Create(ctx context.Context, message *model.Message) error
	Update(ctx context.Context, message *model.Message) error
	Delete(ctx context.Context, id string) error
}
