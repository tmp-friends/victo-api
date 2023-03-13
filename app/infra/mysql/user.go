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
