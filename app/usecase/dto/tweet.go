package dto

import (
	"time"

	"github.com/volatiletech/null/v8"
)

type Tweet struct {
	ID                    string      `boil:"id" json:"id" toml:"id" yaml:"id"`
	Text                  null.String `boil:"text" json:"text,omitempty" toml:"text" yaml:"text,omitempty"`
	RetweetCount          int         `boil:"retweet_count" json:"retweet_count" toml:"retweet_count" yaml:"retweet_count"`
	LikeCount             int         `boil:"like_count" json:"like_count" toml:"like_count" yaml:"like_count"`
	AuthorID              string      `boil:"author_id" json:"author_id" toml:"author_id" yaml:"author_id"`
	URL                   string      `boil:"url" json:"url" toml:"url" yaml:"url"`
	TweetedAt             time.Time   `boil:"tweeted_at" json:"tweeted_at" toml:"tweeted_at" yaml:"tweeted_at"`
	CreatedAt             time.Time   `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	UpdatedAt             time.Time   `boil:"updated_at" json:"updated_at" toml:"updated_at" yaml:"updated_at"`
	HashtagID             int         `boil:"hashtag_id" json:"hashtag_id" toml:"hashtag_id" yaml:"hashtag_id"`
	AuthorName            string      `boil:"author_name" json:"author_name" toml:"author_name" yaml:"author_name"`
	AuthorUsername        string      `boil:"author_username" json:"author_username" toml:"author_username" yaml:"author_username"`
	AuthorProfileImageURL string      `boil:"author_profile_image_url" json:"author_profile_image_url" toml:"author_profile_image_url" yaml:"author_profile_image_url"`
	MediaID               []string    `boil:"media_id" json:"media_id" toml:"media_id" yaml:"media_id"`
	MediaURL              []string    `boil:"media_url" json:"media_url" toml:"media_url" yaml:"media_url"`
	MediaType             []string    `boil:"media_type" json:"media_type" toml:"media_type" yaml:"media_type"`
}
