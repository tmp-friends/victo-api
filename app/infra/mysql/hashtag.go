package mysql

import (
	"context"
	"database/sql"

	"github.com/tmp-friends/victo-api/app/domain/models"
	"github.com/tmp-friends/victo-api/app/domain/repository"
)

type hashtagRepository struct {
	DB *sql.DB
}

func NewHashtagRepository(db *sql.DB) repository.IHashtagRepository {
	return &hashtagRepository{
		DB: db,
	}
}

func (hr *hashtagRepository) FindHashtags(ctx context.Context) (models.HashtagSlice, error) {
	return models.Hashtags().All(ctx, hr.DB)
}
