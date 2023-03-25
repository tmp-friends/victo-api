package mysql

import (
	"context"
	"database/sql"

	"github.com/tmp-friends/victo-api/app/domain/models"
	"github.com/tmp-friends/victo-api/app/usecase/dto"
	"github.com/tmp-friends/victo-api/app/usecase/query"
	"github.com/volatiletech/sqlboiler/v4/boil"
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

func (hq *hashtagQuery) FindHashtag(
	ctx context.Context,
	id int,
	props []string,
	withVtuber bool,
) (dto.Hashtag, error) {
	var hashtag dto.Hashtag

	queries := []qm.QueryMod{}

	queries = append(queries, qm.Where("hashtags.id=?", id))

	switch withVtuber {
	case true:
		queries = append(queries, qm.LeftOuterJoin("vtubers as v on v.id = hashtags.vtuber_id"))
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

	err := models.Hashtags(queries...).Bind(ctx, hq.DB, &hashtag)

	return hashtag, err
}

func (hq *hashtagQuery) FindHashtags(
	ctx context.Context,
	ids []interface{},
	limit int,
	offset int,
	props []string,
	withVtuber bool,
) ([]dto.Hashtag, error) {
	queries := []qm.QueryMod{}

	if ids != nil {
		queries = append(queries, qm.WhereIn("hashtags.id in ?", ids...))
	}
	if limit != 0 {
		queries = append(queries, qm.Limit(limit))
	}
	if offset != 0 {
		queries = append(queries, qm.Offset(offset))
	}
	switch withVtuber {
	case true:
		queries = append(queries, qm.LeftOuterJoin("vtubers as v on v.id = hashtags.vtuber_id"))
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

	var hashtags []dto.Hashtag
	err := models.Hashtags(queries...).Bind(ctx, hq.DB, &hashtags)

	return hashtags, err
}

func (hq *hashtagQuery) FollowHashtag(
	ctx context.Context,
	id int,
	userId int,
) error {
	f := models.HashtagFollow{
		UserID:    userId,
		HashtagID: id,
	}
	err := f.Insert(ctx, hq.DB, boil.Infer())

	return err
}

func (hq *hashtagQuery) UnfollowHashtag(
	ctx context.Context,
	id int,
	userId int,
) error {
	f, _ := models.FindHashtagFollow(ctx, hq.DB, userId, id)
	_, err := f.Delete(ctx, hq.DB)

	return err
}
