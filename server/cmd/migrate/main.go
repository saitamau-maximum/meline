package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/mysqldialect"
	"github.com/uptrace/bun/extra/bundebug"
	"github.com/uptrace/bun/migrate"
	"github.com/urfave/cli/v2"

	"github.com/saitamau-maximum/meline/config"
	infra "github.com/saitamau-maximum/meline/infra/mysql"
	"github.com/saitamau-maximum/meline/migrations"
)

const (
	DRIVER          = "mysql"
	PORT            = "3306"
	NET             = "tcp"
	HOST            = "localhost"
	MIGRATION_TABLE = "bun_migrations"
)

func main() {
	err := config.ValidateDBEnv()
	if err != nil {
		panic(err)
	}
	db, err := infra.ConnectDB(HOST)
	if err != nil {
		panic(err)
	}
	bunDB := bun.NewDB(db, mysqldialect.New())
	bunDB.AddQueryHook(bundebug.NewQueryHook(
		bundebug.WithEnabled(false),
		bundebug.FromEnv(""),
	))

	defer bunDB.Close()

	if err := checkMigrationsTable(context.Background(), bunDB); err != nil {
		log.Println("Migrations table does not exist \n\n\t run `./scripts/migrate.sh init` first")
		return
	}

	app := &cli.App{
		Name: "bun",

		Commands: []*cli.Command{
			newDBCommand(migrate.NewMigrator(bunDB, migration.Migrations)),
		},
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func newDBCommand(migrator *migrate.Migrator) *cli.Command {
	return &cli.Command{
		Name:  "db",
		Usage: "database migration",
		Subcommands: []*cli.Command{
			{
				Name:  "init",
				Usage: "initialize database",
				Action: func(c *cli.Context) error {
					return migrator.Init(c.Context)
				},
			},
			{
				Name:  "migrate",
				Usage: "migrate database",
				Action: func(c *cli.Context) error {
					if err := migrator.Lock(c.Context); err != nil {
						return err
					}
					defer migrator.Unlock(c.Context)

					group, err := migrator.Migrate(c.Context)
					if err != nil {
						return err
					}
					if group.IsZero() {
						fmt.Printf("there are no new migrations to run (database is up to date)\n")
						return nil
					}

					fmt.Printf("migrated %d migrations\n", len(group.Migrations))

					return nil
				},
			},
			{
				Name:  "rollback",
				Usage: "rollback database",
				Action: func(c *cli.Context) error {
					if err := migrator.Lock(c.Context); err != nil {
						return err
					}
					defer migrator.Unlock(c.Context)

					group, err := migrator.Rollback(c.Context)
					if err != nil {
						return err
					}
					if group.IsZero() {
						fmt.Printf("there are no groups to roll back\n")
						return nil
					}

					fmt.Printf("rolled back %d migrations\n", len(group.Migrations))
					return nil
				},
			},
			{
				Name:  "lock",
				Usage: "lock migrations",
				Action: func(c *cli.Context) error {
					return migrator.Lock(c.Context)
				},
			},
			{
				Name:  "unlock",
				Usage: "unlock migrations",
				Action: func(c *cli.Context) error {
					return migrator.Unlock(c.Context)
				},
			},
			{
				Name:  "create_go",
				Usage: "create go migration",
				Action: func(c *cli.Context) error {
					name := strings.Join(c.Args().Slice(), "_")
					mf, err := migrator.CreateGoMigration(c.Context, name)
					if err != nil {
						return err
					}

					fmt.Printf("created migration file %s (%s)\n", mf.Name, mf.Path)
					return nil
				},
			},
			{
				Name:  "create_sql",
				Usage: "create up and down SQL migrations",
				Action: func(c *cli.Context) error {
					name := strings.Join(c.Args().Slice(), "_")
					files, err := migrator.CreateSQLMigrations(c.Context, name)
					if err != nil {
						return err
					}

					for _, mf := range files {
						fmt.Printf("created migration %s (%s)\n", mf.Name, mf.Path)
					}

					return nil
				},
			},
			{
				Name:  "status",
				Usage: "print migrations status",
				Action: func(c *cli.Context) error {
					ms, err := migrator.MigrationsWithStatus(c.Context)
					if err != nil {
						return err
					}
					fmt.Printf("migrations: %s\n", ms)
					fmt.Printf("unapplied migrations: %s\n", ms.Unapplied())
					fmt.Printf("last migration group: %s\n", ms.LastGroup())
					return nil
				},
			},
			{
				Name:  "mark_applied",
				Usage: "mark migrations as applied without actually running them",
				Action: func(c *cli.Context) error {
					group, err := migrator.Migrate(c.Context, migrate.WithNopMigration())
					if err != nil {
						return err
					}
					if group.IsZero() {
						fmt.Printf("there are no new migrations to mark as applied\n")
						return nil
					}
					fmt.Printf("marked as applied %s\n", group)
					return nil
				},
			},
		},
	}
}

func checkMigrationsTable(ctx context.Context, db *bun.DB) error {
	if os.Args[2] == "init" {
		return nil
	}

	if _, err := db.NewSelect().Table(MIGRATION_TABLE).Exists(ctx); err != nil {
		return err
	}

	return nil
}
