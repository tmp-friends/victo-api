package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/tmp-friends/victo-api/app/usecase"
)

type hashtagHandler struct {
	usecase usecase.IHashtagUsecase
}

func NewHashtagHandler(hu usecase.IHashtagUsecase) *hashtagHandler {
	return &hashtagHandler{
		usecase: hu,
	}
}

func (hh *hashtagHandler) FindHashtag() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		p := c.Param("id")
		qms := c.QueryParams()

		hashtag, err := hh.usecase.FindHashtag(ctx, p, qms)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusOK, hashtag)
	}
}

func (hh *hashtagHandler) FindHashtags() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		qms := c.QueryParams()

		hashtags, err := hh.usecase.FindHashtags(ctx, qms)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusOK, hashtags)
	}
}

type FollowHashtagPost struct {
	Id     int `json:"id"`
	UserId int `json:"user_id"`
}

func (hh *hashtagHandler) FollowHashtag() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		p := new(FollowHashtagPost)

		if err := c.Bind(p); err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		err := hh.usecase.FollowHashtag(ctx, p.Id, p.UserId)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusOK, "")
	}
}

func (hh *hashtagHandler) UnfollowHashtag() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		p := new(FollowHashtagPost)

		if err := c.Bind(p); err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		err := hh.usecase.UnfollowHashtag(ctx, p.Id, p.UserId)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusOK, "")
	}
}
