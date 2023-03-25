package usecase

import (
	"context"
	"net/url"
	"strconv"
	"strings"

	"github.com/tmp-friends/victo-api/app/domain/models"
	"github.com/tmp-friends/victo-api/app/usecase/query"
)

type ITweetUsecase interface {
	FindTweet(
		ctx context.Context,
		parameter string,
		qms url.Values,
	) (*models.TweetObject, error)

	FindTweets(
		ctx context.Context,
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
		return nil, err
	}

	props := qms["props"]

	to, err := tu.query.FindTweet(ctx, id, props)
	if err != nil {
		return nil, err
	}

	return to, nil
}

func (tu *tweetUsecase) FindTweets(
	ctx context.Context,
	qms url.Values,
) (models.TweetObjectSlice, error) {
	var hashtagIds []interface{}

	if qms["hashtag_ids"] != nil {
		l := strings.Split(qms["hashtag_ids"][0], ",")

		// sqlboilerでwherein句を使うためにinterface型にする必要あり
		hashtagIds = make([]interface{}, len(l))
		for i, v := range l {
			id, err := strconv.Atoi(v)
			if err != nil {
				return nil, err
			}
			hashtagIds[i] = id
		}
	}

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

	tos, err := tu.query.FindTweets(
		ctx,
		hashtagIds,
		limit,
		offset,
		props,
	)
	if err != nil {
		return nil, err
	}

	return tos, nil
}
