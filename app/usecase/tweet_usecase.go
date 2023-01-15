package usecase

import (
	"context"
	"net/url"

	"github.com/tmp-friends/victo-api/app/domain/models"
	"github.com/tmp-friends/victo-api/app/domain/service"
	"github.com/tmp-friends/victo-api/app/usecase/dto"
)

type ITweetUsecase interface {
	CreateParameter(p string, qms url.Values) dto.FindTweetParameter
	FindTweet(ctx context.Context, parameter dto.FindTweetParameter) (*models.TweetObject, error)
}

type tweetUsecase struct {
	service service.ITweetService
}

func NewTweetUsecase(ts service.ITweetService) ITweetUsecase {
	return &tweetUsecase{
		service: ts,
	}
}

func (tu *tweetUsecase) CreateParameter(p string, qms url.Values) dto.FindTweetParameter {
	return dto.CreateFindTweetParameter(p, qms)
}

func (tu *tweetUsecase) FindTweet(ctx context.Context, parameter dto.FindTweetParameter) (*models.TweetObject, error) {
	mto, err := tu.service.FindTweet(ctx, parameter)
	if err != nil {
		return nil, err
	}

	return mto, nil
}
