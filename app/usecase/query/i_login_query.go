package query

import (
	"context"
	"time"
)

type ILoginQuery interface {
	CreateSessionCookie(
		ctx context.Context,
		idToken string,
	) (string, time.Duration, error)
}
