package repository

import (
	"context"

	"github.com/tmp-friends/victo-api/app/domain/models"
)

type IHashtagRepository interface {
	FindHashtags(ctx context.Context) (models.HashtagSlice, error)
}
