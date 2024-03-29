package usecase

import (
	"context"
	"net/http"
	"net/url"
	"strconv"

	"github.com/tmp-friends/victo-api/app/domain/models"
	"github.com/tmp-friends/victo-api/app/usecase/dto"
	"github.com/tmp-friends/victo-api/app/usecase/query"
)

type IUserUsecase interface {
	GetMyInfo(
		ctx context.Context,
		parameter *http.Cookie,
	) (*models.User, error)

	Login(
		ctx context.Context,
		parameter string,
	) (string, int, error)

	FindFollowingHashtags(
		ctx context.Context,
		parameter string,
		qms url.Values,
	) ([]dto.Hashtag, error)
}

type userUsecase struct {
	mysqlQuery    query.IUserMysqlQuery
	firebaseQuery query.IUserFirebaseQuery
}

func NewUserUsecase(umq query.IUserMysqlQuery, ufq query.IUserFirebaseQuery) IUserUsecase {
	return &userUsecase{
		mysqlQuery:    umq,
		firebaseQuery: ufq,
	}
}

func (uu *userUsecase) GetMyInfo(
	ctx context.Context,
	parameter *http.Cookie,
) (*models.User, error) {
	uid, err := uu.firebaseQuery.GetUIDByCookie(ctx, parameter)
	u, err := uu.mysqlQuery.GetUserByUID(ctx, uid)

	return u, err
}

func (uu *userUsecase) Login(
	ctx context.Context,
	parameter string,
) (string, int, error) {
	googleAcountInfo, err := uu.firebaseQuery.GetGoogleAccountInfoByIdToken(ctx, parameter)
	if err != nil {
		return "", 0, err
	}
	upsertErr := uu.mysqlQuery.UpsertUser(ctx, googleAcountInfo)
	if upsertErr != nil {
		return "", 0, upsertErr
	}

	cookie, expiresIn, err := uu.firebaseQuery.CreateSessionCookie(ctx, parameter)

	return cookie, expiresIn, err
}

func (uu *userUsecase) FindFollowingHashtags(
	ctx context.Context,
	parameter string,
	qms url.Values,
) ([]dto.Hashtag, error) {
	uid, err := strconv.Atoi(parameter)
	if err != nil {
		return []dto.Hashtag{}, err
	}

	props := qms["props"]

	var withVtuber bool
	if qms["withVtuber"] != nil {
		if qms["withVtuber"][0] == "true" {
			withVtuber = true
		}
	}

	return uu.mysqlQuery.FindFollowingHashtags(ctx, uid, props, withVtuber)
}
