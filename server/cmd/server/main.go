package main

import (
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/mysqldialect"

	"github.com/saitamau-maximum/meline/adapter/gateway"
	"github.com/saitamau-maximum/meline/adapter/handler"
	"github.com/saitamau-maximum/meline/adapter/presenter"
	"github.com/saitamau-maximum/meline/config"
	"github.com/saitamau-maximum/meline/domain/entity"
	"github.com/saitamau-maximum/meline/infra/github"
	"github.com/saitamau-maximum/meline/infra/mysql"
	model "github.com/saitamau-maximum/meline/models"
	"github.com/saitamau-maximum/meline/usecase"
)

const (
	HOST = "database"
)

func main() {
	err := config.ValidateAppEnv()
	if err != nil {
		panic(err)
	}

	hub := entity.NewHubEntity()
	go hub.RunLoop()

	e := echo.New()

	db, err := mysql.ConnectDB(HOST)
	if err != nil {
		e.Logger.Error(err)
	}

	bunDB := bun.NewDB(db, mysqldialect.New())
	bunDB.RegisterModel((*model.ChannelUsers)(nil), (*model.ChannelToChannels)(nil), (*model.Channel)(nil), (*model.User)(nil), (*model.MessageToMessages)(nil), (*model.Message)(nil))
	defer bunDB.Close()

	apiGroup := e.Group("/api")

	oAuthConf := github.NewGithubOAuthConf()
	oAuthRepository := github.NewOAuthRepository(oAuthConf)
	userRepository := mysql.NewUserRepository(bunDB)
	channelRepository := mysql.NewChannelRepository(bunDB)
	messageRepository := mysql.NewMessageRepository(bunDB)
	channelUsersRepository := mysql.NewChannelUsersRepository(bunDB)
	channelToChannelsRepository := mysql.NewChannelToChannelsRepository(bunDB)
	githubOAuthInteractor := usecase.NewGithubOAuthInteractor(oAuthRepository)
	authInteractor := usecase.NewAuthInteractor()
	channelInteractor := usecase.NewChannelInteractor(channelRepository, channelUsersRepository, channelToChannelsRepository, userRepository, presenter.NewChannelPresenter())
	userPresenter := presenter.NewUserPresenter()
	userInteractor := usecase.NewUserInteractor(userRepository, userPresenter)
	messageInteractor := usecase.NewMessageInteractor(messageRepository, presenter.NewMessagePresenter())
	clientInteractor := usecase.NewClientInteractor()
	authGateway := gateway.NewAuthGateway(userInteractor)

	handler.NewOAuthHandler(apiGroup.Group("/auth"), githubOAuthInteractor, authInteractor, userInteractor)
	handler.NewUserHandler(apiGroup.Group("/user", authGateway.Auth), userInteractor)
	channelGroup := apiGroup.Group("/channel", authGateway.Auth)
	handler.NewChannelHandler(channelGroup, channelInteractor)
	handler.NewMessageHandler(channelGroup.Group("/:channel_id/message"), messageInteractor, hub)
	handler.NewWebSocketHandler(apiGroup.Group("/ws"), clientInteractor, hub)

	apiGroup.GET("/", authGateway.Auth(func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	}))

	e.Start(":8000")
}
