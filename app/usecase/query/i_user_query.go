package query

import (
	"context"
	"net/http"

	"github.com/tmp-friends/victo-api/app/domain/models"
	"github.com/tmp-friends/victo-api/app/usecase/dto"
)

type IUserMysqlQuery interface {
	GetUserByUID(
		ctx context.Context,
		uid string,
	) (*models.User, error)

	UpsertUser(
		ctx context.Context,
		googleAccountInfo dto.GoogleAccountInfo,
	) error

	FindFollowingHashtags(
		ctx context.Context,
		uid int,
	) (models.HashtagFollowSlice, error)
}

type IUserFirebaseQuery interface {
	GetUIDByCookie(
		ctx context.Context,
		session *http.Cookie,
	) (string, error)

	GetGoogleAccountInfoByIdToken(
		ctx context.Context,
		idToken string,
	) (dto.GoogleAccountInfo, error)

	CreateSessionCookie(
		ctx context.Context,
		idToken string,
	) (string, int, error)
}
