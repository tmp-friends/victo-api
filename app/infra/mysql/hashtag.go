package mysql

import (
	"context"
	"database/sql"

	"github.com/tmp-friends/victo-api/app/domain/models"
	"github.com/tmp-friends/victo-api/app/domain/repository"
	"github.com/tmp-friends/victo-api/app/usecase/dto"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type hashtagRepository struct {
	DB *sql.DB
}

func NewHashtagRepository(db *sql.DB) repository.IHashtagRepository {
	return &hashtagRepository{
		DB: db,
	}
}

func (hr *hashtagRepository) FindHashtags(ctx context.Context, parameter dto.FindHashtagsParameter) (models.HashtagSlice, error) {
	queries := hr.createQueries(parameter)

	return models.Hashtags(queries...).All(ctx, hr.DB)
}

func (hr *hashtagRepository) createQueries(parameter dto.FindHashtagsParameter) []qm.QueryMod {
	queries := []qm.QueryMod{}

	if parameter.Limit != 0 {
		queries = append(queries, qm.Limit(parameter.Limit))
	}

	if parameter.Offset != 0 {
		queries = append(queries, qm.Offset(parameter.Offset))
	}

	if parameter.Props != nil {
		queries = append(queries, qm.Select(parameter.Props...))
	}

	return queries
}
