package query

import (
	"context"

	"github.com/tmp-friends/victo-api/app/usecase/dto"
)

type IHashtagQuery interface {
	FindHashtags(
		ctx context.Context,
		limit int,
		offset int,
		props []string,
		withVtuber bool,
	) ([]dto.Hashtag, error)
}
