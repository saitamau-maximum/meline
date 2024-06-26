package main

import (
	"context"
	"log"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/mysqldialect"

	"github.com/saitamau-maximum/meline/cmd/seeder/seeds"
	"github.com/saitamau-maximum/meline/config"
	infra "github.com/saitamau-maximum/meline/infra/mysql"
)

const (
	HOST = "localhost"
)

func main() {
	err := config.ValidateDBEnv()
	if err != nil {
		panic(err)
	}
	db, err := infra.ConnectDB(HOST)
	if err != nil {
		log.Printf("failed to connect db: %v", err)
		return
	}

	bunDB := bun.NewDB(db, mysqldialect.New())
	defer bunDB.Close()

	if err := seed(context.Background(), bunDB); err != nil {
		log.Printf("failed to seed: %v", err)
		return
	}

	log.Printf("seeding is done")
}

func seed(ctx context.Context, db *bun.DB) error {
	// add seeders here
	if err := seeds.UserSeeds(ctx, db); err != nil {
		return err
	}

	return nil
}
