package mysql

import (
	"context"
	"database/sql"

	"github.com/tmp-friends/victo-api/app/domain/models"
	"github.com/tmp-friends/victo-api/app/domain/repository"
	"github.com/tmp-friends/victo-api/app/usecase/dto"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type tweetRepository struct {
	DB *sql.DB
}

func NewTweetRepository(db *sql.DB) repository.ITweetRepository {
	return &tweetRepository{
		DB: db,
	}
}

// TODO: createQueries()はエンドポイント毎に異なるので、ファイルを分離するか検討
func (tr *tweetRepository) FindTweet(ctx context.Context, parameter dto.FindTweetParameter) (*models.TweetObject, error) {
	queries := tr.createQueries(parameter)

	return models.TweetObjects(queries...).One(ctx, tr.DB)
}

func (tr *tweetRepository) createQueries(parameter dto.FindTweetParameter) []qm.QueryMod {
	queries := []qm.QueryMod{}

	if parameter.Props != nil {
		queries = append(queries, qm.Select(parameter.Props...))
	}

	queries = append(queries, qm.Where("tweetId=?", parameter.Id))

	return queries
}
