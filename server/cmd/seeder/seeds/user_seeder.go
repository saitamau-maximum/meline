package seeds

import (
	"context"
	"errors"

	"github.com/uptrace/bun"
	"github.com/saitamau-maximum/meline/models"
)

var (
	users = []models.User{
		{
			ProviderID: "sample-1",
			Name: "test-user-1",
		},
		{
			ProviderID: "sample-2",
			Name: "test-user-2",
		},
		{
			ProviderID: "sample-3",
			Name: "test-user-3",
		},
	}
)

func UserSeeds(ctx context.Context, db *bun.DB) error {
	if  isExists, err := db.NewSelect().Model(&models.User{}).Exists(ctx); err != nil {
		return err
	} else if isExists {
		return errors.New("user data already exists")
	}

	if _, err := db.NewInsert().Model(&users).Exec(ctx); err != nil {
		return err
	}

	return nil
}
