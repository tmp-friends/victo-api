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
