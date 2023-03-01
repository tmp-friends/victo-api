package mysql

import (
	"context"
	"database/sql"

	"github.com/tmp-friends/victo-api/app/domain/models"
	"github.com/tmp-friends/victo-api/app/usecase/query"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type tweetQuery struct {
	DB *sql.DB
}

func NewTweetQuery(db *sql.DB) query.ITweetQuery {
	return &tweetQuery{
		DB: db,
	}
}

func (tr *tweetQuery) FindTweet(
	ctx context.Context,
	id int,
	props []string,
) (*models.TweetObject, error) {
	queries := []qm.QueryMod{}

	queries = append(queries, qm.Where("id=?", id))

	if props != nil {
		queries = append(queries, qm.Select(props...))
	}

	return models.TweetObjects(queries...).One(ctx, tr.DB)
}

func (tr *tweetQuery) FindTweetsByHashtagId(
	ctx context.Context,
	hashtagId string,
	limit int,
	offset int,
	props []string,
) (models.TweetObjectSlice, error) {
	queries := []qm.QueryMod{}

	queries = append(queries, qm.Where("hashtag_id=?", hashtagId))

	if limit != 0 {
		queries = append(queries, qm.Limit(limit))
	}
	if offset != 0 {
		queries = append(queries, qm.Offset(offset))
	}
	if props != nil {
		queries = append(queries, qm.Select(props...))
	}
	// TODO: sort
	// Enumで指定できるようにする
	queries = append(queries, qm.OrderBy("tweeted_at desc"))

	return models.TweetObjects(queries...).All(ctx, tr.DB)
}
