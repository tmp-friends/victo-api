package query

import (
	"context"

	"github.com/tmp-friends/victo-api/app/domain/models"
	"github.com/tmp-friends/victo-api/app/usecase/dto"
)

type ITweetQuery interface {
	FindTweet(ctx context.Context, parameter dto.FindTweetParameter) (*models.TweetObject, error)
}