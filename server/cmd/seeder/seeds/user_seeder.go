package seeds

import (
	"context"

	"github.com/uptrace/bun"
	"github.com/saitamau-maximum/meline/usecase/model"
)

var (
	users = []model.User{
		{
			GithubID: "sample-1",
			Name: "test-user-1",
		},
		{
			GithubID: "sample-2",
			Name: "test-user-2",
		},
		{
			GithubID: "sample-3",
			Name: "test-user-3",
		},
	}
)

func UserSeeds(ctx context.Context, db *bun.DB) error {
	if _, err := db.NewInsert().Model(&users).Exec(ctx); err != nil {
		return err
	}

	return nil
}
