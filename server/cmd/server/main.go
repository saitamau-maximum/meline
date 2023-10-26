package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/mysqldialect"
	_ "github.com/go-sql-driver/mysql"

	infra "github.com/saitamau-maximum/meline/infra/mysql"
)

const (
	host = "database"
)

func main() {
	e := echo.New()

	db, err := infra.ConnectDB(host)
	if err != nil {
		e.Logger.Error(err)
	}

	bunDB := bun.NewDB(db, mysqldialect.New())
	defer bunDB.Close()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.Start(":8000")
}

