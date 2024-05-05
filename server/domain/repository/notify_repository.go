package repository

import (
	"context"

	model "github.com/saitamau-maximum/meline/models"
)

type INotifyRepository interface {
	FindByUserID(ctx context.Context, userID uint64) ([]*model.Notify, error)
	BulkCreate(ctx context.Context, notifies []model.Notify) error
}
