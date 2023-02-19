package query

import (
	"context"

	"github.com/tmp-friends/victo-api/app/domain/models"
)

type ITweetQuery interface {
	FindTweet(
		ctx context.Context,
		id int,
		props []string,
	) (*models.TweetObject, error)

	FindTweetsByHashtagId(
		ctx context.Context,
		hashtagId string,
		limit int,
		offset int,
		props []string,
	) ([]*models.TweetObject, error)
}
