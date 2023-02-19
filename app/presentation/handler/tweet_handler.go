package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/tmp-friends/victo-api/app/usecase"
)

type tweetHandler struct {
	usecase usecase.ITweetUsecase
}

func NewTweetHandler(tu usecase.ITweetUsecase) *tweetHandler {
	return &tweetHandler{
		usecase: tu,
	}
}

func (th *tweetHandler) FindTweet() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		p := c.Param("id")
		qms := c.QueryParams()

		tweet, err := th.usecase.FindTweet(ctx, p, qms)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusOK, tweet)
	}
}

func (th *tweetHandler) FindTweetsByHashtagId() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		p := c.Param("id")
		qms := c.QueryParams()

		tweets, err := th.usecase.FindTweetsByHashtagId(ctx, p, qms)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusOK, tweets)
	}
}
