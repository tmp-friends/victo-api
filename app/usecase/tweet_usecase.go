package usecase

import (
	"context"
	"net/url"
	"strconv"

	"github.com/tmp-friends/victo-api/app/domain/models"
	"github.com/tmp-friends/victo-api/app/usecase/query"
)

type ITweetUsecase interface {
	FindTweet(
		ctx context.Context,
		parameter string,
		qms url.Values,
	) (*models.TweetObject, error)

	FindTweetsByHashtagId(
		ctx context.Context,
		parameter string,
		qms url.Values,
	) (models.TweetObjectSlice, error)
}

type tweetUsecase struct {
	query query.ITweetQuery
}

func NewTweetUsecase(tq query.ITweetQuery) ITweetUsecase {
	return &tweetUsecase{
		query: tq,
	}
}

func (tu *tweetUsecase) FindTweet(
	ctx context.Context,
	parameter string,
	qms url.Values,
) (*models.TweetObject, error) {
	id, err := strconv.Atoi(parameter)
	if err != nil {
		panic(err)
	}

	props := qms["props"]

	to, err := tu.query.FindTweet(ctx, id, props)
	if err != nil {
		return nil, err
	}

	return to, nil
}

func (tu *tweetUsecase) FindTweetsByHashtagId(
	ctx context.Context,
	parameter string,
	qms url.Values,
) (models.TweetObjectSlice, error) {
	hashtagId := parameter

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

	tos, err := tu.query.FindTweetsByHashtagId(
		ctx,
		hashtagId,
		limit,
		offset,
		props,
	)
	if err != nil {
		return nil, err
	}

	return tos, nil
}
