package migration

import (
	"context"
	"fmt"

	"github.com/uptrace/bun"
	"github.com/saitamau-maximum/meline/models"
)

func init() {
	Migrations.MustRegister(func(ctx context.Context, db *bun.DB) error {
		fmt.Print(" [up migration] ")
		db.NewCreateTable().Model((*model.User)(nil)).Exec(ctx)
		return nil
	}, func(ctx context.Context, db *bun.DB) error {
		fmt.Print(" [down migration] ")
		db.NewDropTable().Model((*model.User)(nil)).IfExists().Exec(ctx)
		return nil
	})
}
