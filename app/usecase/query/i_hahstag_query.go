package query

import (
	"context"

	"github.com/tmp-friends/victo-api/app/usecase/dto"
)

type IHashtagQuery interface {
	FindHashtag(
		ctx context.Context,
		id int,
		props []string,
		withVtuber bool,
	) (dto.Hashtag, error)

	FindHashtags(
		ctx context.Context,
		limit int,
		offset int,
		props []string,
		withVtuber bool,
	) ([]dto.Hashtag, error)

	FollowHashtag(
		ctx context.Context,
		id int,
		userId int,
	) error
}
