package mysql

import (
	"context"
	"database/sql"
	"strings"
	"time"

	"github.com/tmp-friends/victo-api/app/domain/models"
	"github.com/tmp-friends/victo-api/app/usecase/dto"
	"github.com/tmp-friends/victo-api/app/usecase/query"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type tweetQuery struct {
	DB *sql.DB
}

func NewTweetQuery(db *sql.DB) query.ITweetQuery {
	return &tweetQuery{
		DB: db,
	}
}

type TweetRawData struct {
	ID                    string      `boil:"id"`
	Text                  null.String `boil:"text"`
	RetweetCount          int         `boil:"retweet_count"`
	LikeCount             int         `boil:"like_count"`
	AuthorID              string      `boil:"author_id"`
	URL                   string      `boil:"url"`
	TweetedAt             time.Time   `boil:"tweeted_at"`
	CreatedAt             time.Time   `boil:"created_at"`
	UpdatedAt             time.Time   `boil:"updated_at"`
	HashtagID             int         `boil:"hashtag_id"`
	AuthorName            string      `boil:"author_name"`
	AuthorUsername        string      `boil:"author_username"`
	AuthorProfileImageURL string      `boil:"author_profile_image_url"`
	MediaID               string      `boil:"media_id"`
	MediaURL              string      `boil:"media_url"`
	MediaType             string      `boil:"media_type"`
}

func (tr *tweetQuery) FindTweet(
	ctx context.Context,
	id int,
	props []string,
) (*models.TweetObject, error) {
	queries := []qm.QueryMod{}

	queries = append(queries, qm.Where("id=?", id))

	if props != nil {
		queries = append(queries, qm.Select(props...))
	}

	return models.TweetObjects(queries...).One(ctx, tr.DB)
}

func (tr *tweetQuery) FindTweets(
	ctx context.Context,
	hashtagIds []interface{},
	limit int,
	offset int,
	props []string,
	withMedia bool,
) ([]dto.Tweet, error) {
	var tweetsRawData []TweetRawData

	queries := []qm.QueryMod{}

	// Authorsテーブルとjoin
	queries = append(queries, qm.LeftOuterJoin("authors as a on tweet_objects.author_id = a.id"))
	queries = append(queries, qm.Select(
		"tweet_objects.*",
		"a.name as author_name",
		"a.username as author_username",
		"a.profile_image_url as author_profile_image_url",
	))

	if hashtagIds != nil {
		queries = append(queries, qm.WhereIn("hashtag_id in ?", hashtagIds...))
	}

	if withMedia {
		queries = append(queries, qm.LeftOuterJoin("media_objects as m on tweet_objects.id = m.tweet_id"))

		queries = append(queries, qm.Select(
			"GROUP_CONCAT(m.id) as media_id",
			"GROUP_CONCAT(m.url) as media_url",
			"GROUP_CONCAT(m.type) as media_type",
		))
		queries = append(queries, qm.GroupBy(
			"id",
		))
	}

	if limit != 0 {
		queries = append(queries, qm.Limit(limit))
	}
	if offset != 0 {
		queries = append(queries, qm.Offset(offset))
	}
	if props != nil {
		queries = append(queries, qm.Select(props...))
	}

	// TODO: sort
	// Enumで指定できるようにする
	queries = append(queries, qm.OrderBy("tweeted_at desc"))

	err := models.TweetObjects(queries...).Bind(ctx, tr.DB, &tweetsRawData)

	return tr.createDto(tweetsRawData), err
}

func (tr *tweetQuery) createDto(
	rawData []TweetRawData,
) []dto.Tweet {
	tweets := make([]dto.Tweet, len(rawData))

	for i, v := range rawData {
		tweets[i].ID = v.ID
		tweets[i].Text = v.Text
		tweets[i].RetweetCount = v.RetweetCount
		tweets[i].LikeCount = v.LikeCount
		tweets[i].AuthorID = v.AuthorID
		tweets[i].URL = v.URL
		tweets[i].TweetedAt = v.TweetedAt
		tweets[i].CreatedAt = v.CreatedAt
		tweets[i].UpdatedAt = v.UpdatedAt
		tweets[i].HashtagID = v.HashtagID
		tweets[i].AuthorName = v.AuthorName
		tweets[i].AuthorUsername = v.AuthorUsername
		tweets[i].AuthorProfileImageURL = v.AuthorProfileImageURL
		tweets[i].MediaID = strings.Split(v.MediaID, ",")
		tweets[i].MediaURL = strings.Split(v.MediaURL, ",")
		tweets[i].MediaType = strings.Split(v.MediaType, ",")
	}

	return tweets
}
