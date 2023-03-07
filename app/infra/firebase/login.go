package firebase

import (
	"context"
	"time"

	firebase "firebase.google.com/go/v4"
	"github.com/tmp-friends/victo-api/app/usecase/query"
)

type loginQuery struct {
	Firebase *firebase.App
}

func NewLoginQuery(firebase *firebase.App) query.ILoginQuery {
	return &loginQuery{
		Firebase: firebase,
	}
}

// @see: https://firebase.google.com/docs/auth/admin/manage-cookies?hl=ja
func (fq *loginQuery) CreateSessionCookie(
	ctx context.Context,
	idToken string,
) (string, time.Duration, error) {
	client, _ := fq.Firebase.Auth(ctx)

	expiresIn := time.Hour * 24 * 14

	cookie, err := client.SessionCookie(ctx, idToken, expiresIn)

	return cookie, expiresIn, err
}
