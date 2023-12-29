package migration

import (
	"context"
	"fmt"

	model "github.com/saitamau-maximum/meline/models"
	"github.com/uptrace/bun"
)

func init() {
	Migrations.MustRegister(func(ctx context.Context, db *bun.DB) error {
		fmt.Print(" [up migration] ")
		db.NewCreateTable().Model((*model.ChannelUsers)(nil)).Exec(ctx)
		return nil
	}, func(ctx context.Context, db *bun.DB) error {
		fmt.Print(" [down migration] ")
		db.NewDropTable().Model((*model.ChannelUsers)(nil)).IfExists().Exec(ctx)
		return nil
	})
}
