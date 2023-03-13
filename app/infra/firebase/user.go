package firebase

import (
	"context"
	"net/http"
	"time"

	firebase "firebase.google.com/go/v4"
	"github.com/tmp-friends/victo-api/app/usecase/dto"
	"github.com/tmp-friends/victo-api/app/usecase/query"
)

type userQuery struct {
	Firebase *firebase.App
}

func NewUserQuery(firebase *firebase.App) query.IUserFirebaseQuery {
	return &userQuery{
		Firebase: firebase,
	}
}

func (uq *userQuery) GetUIDByCookie(
	ctx context.Context,
	session *http.Cookie,
) (string, error) {
	client, _ := uq.Firebase.Auth(ctx)

	decoded, err := client.VerifySessionCookie(ctx, session.Value)
	if err != nil {
		return "", err
	}

	uid := decoded.UID

	return uid, err
}

func (uq *userQuery) GetGoogleAccountInfoByIdToken(
	ctx context.Context,
	idToken string,
) (dto.GoogleAccountInfo, error) {
	client, _ := uq.Firebase.Auth(ctx)

	decoded, err := client.VerifyIDToken(ctx, idToken)
	if err != nil {
		return dto.GoogleAccountInfo{}, err
	}

	googleAccountInfo := dto.GoogleAccountInfo{
		UID:     decoded.UID,
		Email:   decoded.Claims["email"].(string),
		Name:    decoded.Claims["name"].(string),
		Picture: decoded.Claims["picture"].(string),
	}

	return googleAccountInfo, err
}

// @see: https://firebase.google.com/docs/auth/admin/manage-cookies?hl=ja
func (uq *userQuery) CreateSessionCookie(
	ctx context.Context,
	idToken string,
) (string, int, error) {
	client, _ := uq.Firebase.Auth(ctx)

	expiresIn := time.Hour * 24 * 14

	cookie, err := client.SessionCookie(ctx, idToken, expiresIn)

	return cookie, int(expiresIn.Seconds()), err
}
