package mysql

import (
	"context"
	"database/sql"

	"github.com/tmp-friends/victo-api/app/domain/models"
	"github.com/tmp-friends/victo-api/app/usecase/query"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type hashtagQuery struct {
	DB *sql.DB
}

func NewHashtagQuery(db *sql.DB) query.IHashtagQuery {
	return &hashtagQuery{
		DB: db,
	}
}

func (hr *hashtagQuery) FindHashtags(
	ctx context.Context,
	limit int,
	offset int,
	props []string,
) (models.HashtagSlice, error) {
	queries := []qm.QueryMod{}

	if limit != 0 {
		queries = append(queries, qm.Limit(limit))
	}
	if offset != 0 {
		queries = append(queries, qm.Offset(offset))
	}
	if props != nil {
		queries = append(queries, qm.Select(props...))
	}

	return models.Hashtags(queries...).All(ctx, hr.DB)
}
