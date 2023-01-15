package usecase

import (
	"context"
	"net/url"

	"github.com/tmp-friends/victo-api/app/domain/models"
	"github.com/tmp-friends/victo-api/app/usecase/dto"
	"github.com/tmp-friends/victo-api/app/usecase/query"
)

type ITweetUsecase interface {
	CreateParameter(p string, qms url.Values) dto.FindTweetParameter
	FindTweet(ctx context.Context, parameter dto.FindTweetParameter) (*models.TweetObject, error)
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
