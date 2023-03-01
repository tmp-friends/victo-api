package usecase

import (
	"context"
	"net/url"
	"strconv"

	"github.com/tmp-friends/victo-api/app/usecase/dto"
	"github.com/tmp-friends/victo-api/app/usecase/query"
)

type IHashtagUsecase interface {
	FindHashtags(ctx context.Context, qms url.Values) ([]dto.Hashtag, error)
}

type hashtagUsecase struct {
	query query.IHashtagQuery
}

func NewHashtagUsecase(hq query.IHashtagQuery) IHashtagUsecase {
	return &hashtagUsecase{
		query: hq,
	}
}

func (hu *hashtagUsecase) FindHashtags(
	ctx context.Context,
	qms url.Values,
) ([]dto.Hashtag, error) {
	// TODO: createParameters()を作るか検討
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

	hs, err := hu.query.FindHashtags(ctx, limit, offset, props, withVtuber)
	if err != nil {
		return nil, err
	}

	return hs, nil
}
