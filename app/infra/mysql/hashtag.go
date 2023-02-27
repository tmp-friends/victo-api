package mysql

import (
	"context"
	"database/sql"

	"github.com/tmp-friends/victo-api/app/domain/models"
	"github.com/tmp-friends/victo-api/app/usecase/dto"
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
	withVtuber bool,
) ([]dto.Hashtag, error) {
	var hashtags []dto.Hashtag

	queries := []qm.QueryMod{}

	if limit != 0 {
		queries = append(queries, qm.Limit(limit))
	}
	if offset != 0 {
		queries = append(queries, qm.Offset(offset))
	}
	switch withVtuber {
	case true:
		queries = append(queries, qm.LeftOuterJoin("vtubers as v on v.ID = hashtags.vtuber_id"))
		// TODO: props指定
		queries = append(queries, qm.Select(
			"hashtags.*",
			"v.name as vtuber_name",
			"v.belongs_to as belongs_to",
			"v.profile_image_url as profile_image_url",
			"v.twitter_user_name as twitter_user_name",
			"v.channel as channel",
		))
	default:
		if props != nil {
			queries = append(queries, qm.Select(props...))
		}
	}

	err := models.Hashtags(queries...).Bind(ctx, hr.DB, &hashtags)

	return hashtags, err
}
