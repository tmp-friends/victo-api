package query

import (
	"context"

	"github.com/tmp-friends/victo-api/app/domain/models"
	"github.com/tmp-friends/victo-api/app/usecase/dto"
)

type ITweetQuery interface {
	FindTweet(
		ctx context.Context,
		id int,
		props []string,
	) (*models.TweetObject, error)

	FindTweets(
		ctx context.Context,
		hashtagIds []interface{},
		limit int,
		offset int,
		props []string,
		withMedia bool,
	) ([]dto.Tweet, error)
}
