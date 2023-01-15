package config

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/tmp-friends/victo-api/app/domain/service"
	"github.com/tmp-friends/victo-api/app/infra/mysql"
	"github.com/tmp-friends/victo-api/app/presentation/handler"
	"github.com/tmp-friends/victo-api/app/usecase"
)

const (
	apiVersion = "/v1"
)

func InitRouter() *echo.Echo {
	e := echo.New()

	e.Use(
		middleware.Logger(),
		middleware.Recover(),
		middleware.CORS(),
	)

	mysqlConnector := NewMySQLConnector()

	// health check
	healthzGroup := e.Group("/healthz")
	{
		healthzGroup.GET("", handler.Check())
	}

	// hashtag
	hashtagRepository := mysql.NewHashtagRepository(mysqlConnector.Conn)
	hashtagService := service.NewHashtagService(hashtagRepository)
	hashtagUsecase := usecase.NewHashtagUsecase(hashtagService)

	hashtagsGroup := e.Group("/hashtags")
	{
		hashtagHandler := handler.NewHashtagHandler(hashtagUsecase)
		hashtagsGroup.GET("", hashtagHandler.FindHashtags())
	}

	// tweet
	tweetRepository := mysql.NewTweetRepository(mysqlConnector.Conn)
	tweetService := service.NewTweetService(tweetRepository)
	tweetUsecase := usecase.NewTweetUsecase(tweetService)

	tweetGroup := e.Group("/tweet")
	{
		tweetHandler := handler.NewTweetHandler(tweetUsecase)
		tweetGroup.GET("/:id", tweetHandler.FindTweet())
	}

	return e
}
