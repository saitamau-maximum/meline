package main

import (
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/mysqldialect"

	"github.com/saitamau-maximum/meline/controller/handler"
	"github.com/saitamau-maximum/meline/controller/gateway"
	"github.com/saitamau-maximum/meline/infra/github"
	"github.com/saitamau-maximum/meline/infra/mysql"
	"github.com/saitamau-maximum/meline/usecase"
)

const (
	HOST = "database"
)

func main() {
	e := echo.New()

	db, err := mysql.ConnectDB(HOST)
	if err != nil {
		e.Logger.Error(err)
	}

	bunDB := bun.NewDB(db, mysqldialect.New())
	defer bunDB.Close()

	apiGroup := e.Group("/api")

	oAuthConf := github.NewGithubOAuthConf()
	authRepository := github.NewAuthRepository(oAuthConf)
	userRepository := mysql.NewUserRepository(bunDB)
	authInteractor := usecase.NewAuthInteractor(authRepository)
	userInteractor := usecase.NewUserInteractor(userRepository)

	requiredAuthGroup := gateway.NewAuthGateway(apiGroup, userInteractor)
	
	handler.NewAuthHandler(apiGroup, authInteractor, userInteractor)

	requiredAuthGroup.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.Start(":8000")
}

