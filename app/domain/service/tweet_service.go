package service

import (
	"context"

	"github.com/tmp-friends/victo-api/app/domain/models"
	"github.com/tmp-friends/victo-api/app/domain/repository"
	"github.com/tmp-friends/victo-api/app/usecase/dto"
)

type ITweetService interface {
	FindTweet(ctx context.Context, parameter dto.FindTweetParameter) (*models.TweetObject, error)
}

type tweetService struct {
	repo repository.ITweetRepository
}

func NewTweetService(tr repository.ITweetRepository) ITweetService {
	return &tweetService{
		repo: tr,
	}
}

func (ts *tweetService) FindTweet(ctx context.Context, parameter dto.FindTweetParameter) (*models.TweetObject, error) {
	return ts.repo.FindTweet(ctx, parameter)
}
