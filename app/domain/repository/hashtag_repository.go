package repository

import (
	"context"

	"github.com/tmp-friends/victo-api/app/domain/models"
	"github.com/tmp-friends/victo-api/app/usecase/dto"
)

type IHashtagRepository interface {
	FindHashtags(ctx context.Context, parameter dto.FindHashtagsParameter) (models.HashtagSlice, error)
}
