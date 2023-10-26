package main

import (
	"net/http"
	"context"

	"github.com/labstack/echo/v4"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/mysqldialect"
	_ "github.com/go-sql-driver/mysql"

	"github.com/saitamau-maximum/meline/usecase/model"
	infra "github.com/saitamau-maximum/meline/infra/mysql"
)

func main() {
	e := echo.New()

	db, err := infra.ConnectDB()
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

