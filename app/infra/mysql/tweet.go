package mysql

import (
	"context"
	"database/sql"

	"github.com/tmp-friends/victo-api/app/domain/models"
	"github.com/tmp-friends/victo-api/app/usecase/dto"
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

// TODO: createQueries()はエンドポイント毎に異なるので、ファイルを分離するか検討
func (tr *tweetQuery) FindTweet(ctx context.Context, parameter dto.FindTweetParameter) (*models.TweetObject, error) {
	queries := tr.createQueries(parameter)

	return models.TweetObjects(queries...).One(ctx, tr.DB)
}

func (tr *tweetQuery) createQueries(parameter dto.FindTweetParameter) []qm.QueryMod {
	queries := []qm.QueryMod{}

	if parameter.Props != nil {
		queries = append(queries, qm.Select(parameter.Props...))
	}

	queries = append(queries, qm.Where("tweetId=?", parameter.Id))

	return queries
}
