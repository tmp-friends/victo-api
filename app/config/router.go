package config

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/tmp-friends/victo-api/app/infra/mysql"
	"github.com/tmp-friends/victo-api/app/presentation/handler"
	"github.com/tmp-friends/victo-api/app/usecase"
)

const (
	apiVersion   = "/v1"
	healthzRoot  = "/healthz"
	hashtagsRoot = apiVersion + "/hashtags"
	tweetsRoot   = apiVersion + "/tweets"
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
	healthzGroup := e.Group(healthzRoot)
	{
		healthzGroup.GET("", handler.Check())
	}

	// hashtag
	hashtagQuery := mysql.NewHashtagQuery(mysqlConnector.Conn)
	hashtagUsecase := usecase.NewHashtagUsecase(hashtagQuery)

	hashtagsGroup := e.Group(hashtagsRoot)
	{
		hashtagHandler := handler.NewHashtagHandler(hashtagUsecase)
		hashtagsGroup.GET("/:id", hashtagHandler.FindHashtag())
		hashtagsGroup.GET("", hashtagHandler.FindHashtags())
	}

	// tweet
	tweetQuery := mysql.NewTweetQuery(mysqlConnector.Conn)
	tweetUsecase := usecase.NewTweetUsecase(tweetQuery)

	tweetsGroup := e.Group(tweetsRoot)
	{
		tweetHandler := handler.NewTweetHandler(tweetUsecase)
		tweetsGroup.GET("/:id", tweetHandler.FindTweet())
		tweetsGroup.GET("/hashtag/:id", tweetHandler.FindTweetsByHashtagId())
	}

	return e
}
