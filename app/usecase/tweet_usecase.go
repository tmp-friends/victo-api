package usecase

import (
	"context"
	"net/url"
	"strconv"

	"github.com/tmp-friends/victo-api/app/domain/models"
	"github.com/tmp-friends/victo-api/app/usecase/dto"
	"github.com/tmp-friends/victo-api/app/usecase/query"
)

type ITweetUsecase interface {
	CreateParameter(p string, qms url.Values) dto.FindTweetParameter
	FindTweet(ctx context.Context, parameter dto.FindTweetParameter) (*models.TweetObject, error)

	FindTweetsByHashtagId(
		ctx context.Context,
		parameter string,
		qms url.Values,
	) ([]*models.TweetObject, error)
}

type tweetUsecase struct {
	query query.ITweetQuery
}

func NewTweetUsecase(tq query.ITweetQuery) ITweetUsecase {
	return &tweetUsecase{
		query: tq,
	}
}

func (tu *tweetUsecase) CreateParameter(p string, qms url.Values) dto.FindTweetParameter {
	return dto.CreateFindTweetParameter(p, qms)
}

func (tu *tweetUsecase) FindTweet(ctx context.Context, parameter dto.FindTweetParameter) (*models.TweetObject, error) {
	mto, err := tu.query.FindTweet(ctx, parameter)
	if err != nil {
		return nil, err
	}

	return mto, nil
}

func (tu *tweetUsecase) FindTweetsByHashtagId(
	ctx context.Context,
	parameter string,
	qms url.Values,
) ([]*models.TweetObject, error) {
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

	mto, err := tu.query.FindTweetsByHashtagId(
		ctx,
		hashtagId,
		limit,
		offset,
		props,
	)
	if err != nil {
		return nil, err
	}

	return mto, nil
}
