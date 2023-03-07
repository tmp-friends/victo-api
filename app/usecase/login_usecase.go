package usecase

import (
	"context"
	"time"

	"github.com/tmp-friends/victo-api/app/usecase/query"
)

type ILoginUsecase interface {
	CreateSessionCookie(
		ctx context.Context,
		parameter string,
	) (string, time.Duration, error)
}

type loginUsecase struct {
	query query.ILoginQuery
}

func NewLoginUsecase(lq query.ILoginQuery) ILoginUsecase {
	return &loginUsecase{
		query: lq,
	}
}

func (lu *loginUsecase) CreateSessionCookie(
	ctx context.Context,
	parameter string,
) (string, time.Duration, error) {
	cookie, expiresIn, err := lu.query.CreateSessionCookie(ctx, parameter)

	return cookie, expiresIn, err
}
