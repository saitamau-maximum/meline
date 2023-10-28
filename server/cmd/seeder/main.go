package main

import (
	"context"
	"log"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/mysqldialect"
	infra "github.com/saitamau-maximum/meline/infra/mysql"
	model "github.com/saitamau-maximum/meline/usecase/model"
)

const (
	HOST = "localhost"
)

var (
	users = []*model.User{
		{
			GithubID: "sample1",
			Name: "name_1",
		},
		{
			GithubID: "sample2",
			Name: "name_2",
		},
		{
			GithubID: "sample3",
			Name: "name_3",
		},
	}		
)

func main() {
	db, err := infra.ConnectDB(HOST)
	if err != nil {
		log.Printf("failed to connect db: %v", err)
		return
	}

	bunDB := bun.NewDB(db, mysqldialect.New())
	defer bunDB.Close()

	if err := seeds(context.Background(), bunDB); err != nil {
		log.Printf("failed to seed: %v", err)
		return
	}

	log.Printf("seeding is done")
}

func seeds(ctx context.Context, db *bun.DB) error {
	if _, err := db.NewInsert().Model(&users).Exec(ctx); err != nil {
		return err
	}

	return nil
}
