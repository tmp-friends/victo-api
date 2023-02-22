package query

import (
	"context"

	"github.com/tmp-friends/victo-api/app/domain/models"
)

type IHashtagQuery interface {
	FindHashtags(
		ctx context.Context,
		limit int,
		offset int,
		props []string,
	) (models.HashtagSlice, error)
}
