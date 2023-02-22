package usecase

import (
	"context"
	"net/url"
	"strconv"

	"github.com/tmp-friends/victo-api/app/domain/models"
	"github.com/tmp-friends/victo-api/app/usecase/query"
)

type IHashtagUsecase interface {
	FindHashtags(ctx context.Context, qms url.Values) (models.HashtagSlice, error)
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
) (models.HashtagSlice, error) {
	var limit int
	if qms["limit"] != nil {
		l, err := strconv.Atoi(qms["limit"][0])
		if err != nil {
			panic(err)
		}
		limit = l
	}

	var offset int
	if qms["offset"] != nil {
		o, err := strconv.Atoi(qms["offset"][0])
		if err != nil {
			panic(err)
		}
		offset = o
	}

	props := qms["props"]

	hs, err := hu.query.FindHashtags(ctx, limit, offset, props)
	if err != nil {
		return nil, err
	}

	return hs, nil
}
