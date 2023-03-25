package config

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/tmp-friends/victo-api/app/infra/firebase"
	"github.com/tmp-friends/victo-api/app/infra/mysql"
	"github.com/tmp-friends/victo-api/app/presentation/handler"
	"github.com/tmp-friends/victo-api/app/usecase"
)

const (
	apiVersion   = "/v1"
	healthzRoot  = "/healthz"
	usersRoot    = apiVersion + "/users"
	hashtagsRoot = apiVersion + "/hashtags"
	tweetsRoot   = apiVersion + "/tweets"
)

func InitRouter() *echo.Echo {
	e := echo.New()

	e.Use(
		middleware.Logger(),
		middleware.Recover(),
		middleware.CORSWithConfig(middleware.CORSConfig{
			AllowOrigins:     []string{"http://localhost:3000"},
			AllowCredentials: true,
		}),
	)

	// create instances
	firebaseApp := NewFirebaseApp()
	mysqlConnector := NewMySQLConnector()

	// health check
	healthzGroup := e.Group(healthzRoot)
	{
		healthzGroup.GET("", handler.Check())
	}

	// user
	userMysqlQuery := mysql.NewUserQuery(mysqlConnector.Conn)
	userFirebaseQuery := firebase.NewUserQuery(firebaseApp)
	userUsecase := usecase.NewUserUsecase(userMysqlQuery, userFirebaseQuery)

	usersGroup := e.Group(usersRoot)
	{
		userHandler := handler.NewUserHandler(userUsecase)
		usersGroup.GET("/me", userHandler.GetMyInfo())
		usersGroup.POST("/login", userHandler.Login())
		// TODO: "/logout"エンドポイントを実装する
		usersGroup.GET("/:id/following_hashtags", userHandler.FindFollowingHashtags())
	}

	// hashtag
	hashtagQuery := mysql.NewHashtagQuery(mysqlConnector.Conn)
	hashtagUsecase := usecase.NewHashtagUsecase(hashtagQuery)

	hashtagsGroup := e.Group(hashtagsRoot)
	{
		hashtagHandler := handler.NewHashtagHandler(hashtagUsecase)
		hashtagsGroup.GET("/:id", hashtagHandler.FindHashtag())
		hashtagsGroup.GET("", hashtagHandler.FindHashtags())
		hashtagsGroup.POST("/follow", hashtagHandler.FollowHashtag())
		hashtagsGroup.POST("/unfollow", hashtagHandler.UnfollowHashtag())
	}

	// tweet
	tweetQuery := mysql.NewTweetQuery(mysqlConnector.Conn)
	tweetUsecase := usecase.NewTweetUsecase(tweetQuery)

	tweetsGroup := e.Group(tweetsRoot)
	{
		tweetHandler := handler.NewTweetHandler(tweetUsecase)
		tweetsGroup.GET("/:id", tweetHandler.FindTweet())
		tweetsGroup.GET("", tweetHandler.FindTweets())
	}

	return e
}
