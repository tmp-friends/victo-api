package mysql

import (
	"context"
	"crypto/sha256"
	"database/sql"
	"encoding/hex"

	"github.com/tmp-friends/victo-api/app/domain/models"
	"github.com/tmp-friends/victo-api/app/usecase/dto"
	"github.com/tmp-friends/victo-api/app/usecase/query"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type userQuery struct {
	DB *sql.DB
}

func NewUserQuery(db *sql.DB) query.IUserMysqlQuery {
	return &userQuery{
		DB: db,
	}
}

func (uq *userQuery) GetUserByUID(
	ctx context.Context,
	uid string,
) (*models.User, error) {
	u, err := models.Users(
		models.UserWhere.HashUID.EQ(uq.hash(uid)),
	).One(ctx, uq.DB)

	return u, err
}

func (uq *userQuery) UpsertUser(
	ctx context.Context,
	googleAcountInfo dto.GoogleAccountInfo,
) error {
	u := models.User{
		HashUID:         uq.hash(googleAcountInfo.UID),
		Name:            googleAcountInfo.Name,
		Email:           googleAcountInfo.Email,
		ProfileImageURL: googleAcountInfo.Picture,
	}

	err := u.Upsert(ctx, uq.DB, boil.Blacklist("created_at"), boil.Infer())

	return err
}

func (uq *userQuery) hash(uid string) string {
	bytes := sha256.Sum256([]byte(uid))

	return hex.EncodeToString(bytes[:])
}

func (uq *userQuery) FindFollowingHashtags(
	ctx context.Context,
	uid int,
	props []string,
	withVtuber bool,
) ([]dto.Hashtag, error) {
	// FollowしているHashtagを取得
	hfs, err := models.HashtagFollows(
		models.HashtagFollowWhere.UserID.EQ(uid),
	).All(ctx, uq.DB)

	// Hashtag情報を取得
	var hashtagIds []interface{}
	for _, v := range hfs {
		hashtagIds = append(hashtagIds, v.HashtagID)
	}

	queries := []qm.QueryMod{}

	queries = append(queries, qm.WhereIn("hashtags.id in ?", hashtagIds...))

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
	err = models.Hashtags(queries...).Bind(ctx, uq.DB, &hashtags)

	return hashtags, err
}
