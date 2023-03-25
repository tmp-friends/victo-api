package usecase

import (
	"context"
	"net/url"
	"strconv"
	"strings"

	"github.com/tmp-friends/victo-api/app/usecase/dto"
	"github.com/tmp-friends/victo-api/app/usecase/query"
)

type IHashtagUsecase interface {
	FindHashtag(
		ctx context.Context,
		parameter string,
		qms url.Values,
	) (dto.Hashtag, error)

	FindHashtags(ctx context.Context, qms url.Values) ([]dto.Hashtag, error)

	FollowHashtag(ctx context.Context, hashtagId int, userId int) error

	UnfollowHashtag(ctx context.Context, hashtagId int, userId int) error
}

type hashtagUsecase struct {
	query query.IHashtagQuery
}

func NewHashtagUsecase(hq query.IHashtagQuery) IHashtagUsecase {
	return &hashtagUsecase{
		query: hq,
	}
}

func (hu *hashtagUsecase) FindHashtag(
	ctx context.Context,
	parameter string,
	qms url.Values,
) (dto.Hashtag, error) {
	id, err := strconv.Atoi(parameter)
	if err != nil {
		return dto.Hashtag{}, err
	}

	props := qms["props"]

	var withVtuber bool
	if qms["withVtuber"] != nil {
		if qms["withVtuber"][0] == "true" {
			withVtuber = true
		}
	}

	h, err := hu.query.FindHashtag(ctx, id, props, withVtuber)
	if err != nil {
		return dto.Hashtag{}, err
	}

	return h, nil
}

func (hu *hashtagUsecase) FindHashtags(
	ctx context.Context,
	qms url.Values,
) ([]dto.Hashtag, error) {
	var ids []interface{}

	if qms["ids"] != nil {
		l := strings.Split(qms["ids"][0], ",")

		// sqlboilerでwherein句を使うためにinterface型にする必要あり
		ids = make([]interface{}, len(l))
		for i, v := range l {
			id, err := strconv.Atoi(v)
			if err != nil {
				return nil, err
			}
			ids[i] = id
		}
	}

	var limit int
	if qms["limit"] != nil {
		l, err := strconv.Atoi(qms["limit"][0])
		if err != nil {
			return nil, err
		}
		limit = l
	}

	var offset int
	if qms["offset"] != nil {
		o, err := strconv.Atoi(qms["offset"][0])
		if err != nil {
			return nil, err
		}
		offset = o
	}

	props := qms["props"]

	var withVtuber bool
	if qms["withVtuber"] != nil {
		if qms["withVtuber"][0] == "true" {
			withVtuber = true
		}
	}

	hs, err := hu.query.FindHashtags(ctx, ids, limit, offset, props, withVtuber)
	if err != nil {
		return nil, err
	}

	return hs, nil
}

func (hu *hashtagUsecase) FollowHashtag(
	ctx context.Context,
	id int,
	userId int,
) error {
	return hu.query.FollowHashtag(ctx, id, userId)
}

func (hu *hashtagUsecase) UnfollowHashtag(
	ctx context.Context,
	id int,
	userId int,
) error {
	return hu.query.UnfollowHashtag(ctx, id, userId)
}
