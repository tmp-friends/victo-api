package dto

import (
	"time"

	"github.com/volatiletech/null/v8"
)

type Hashtag struct {
	ID              int         `boil:"id" json:"id" toml:"id" yaml:"id"`
	Name            string      `boil:"name" json:"name" toml:"name" yaml:"name"`
	IsSelf          bool        `boil:"is_self" json:"is_self" toml:"is_self" yaml:"is_self"`
	CreatedAt       time.Time   `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	UpdatedAt       time.Time   `boil:"updated_at" json:"updated_at" toml:"updated_at" yaml:"updated_at"`
	VtuberID        int         `boil:"vtuber_id" json:"vtuber_id" toml:"vtuber_id" yaml:"vtuber_id"`
	VtuberName      string      `boil:"vtuber_name" json:"vtuber_name" toml:"vtuber_name" yaml:"vtuber_name"`
	BelongsTo       null.String `boil:"belongs_to" json:"belongs_to,omitempty" toml:"belongs_to" yaml:"belongs_to,omitempty"`
	ProfileImageURL null.String `boil:"profile_image_url" json:"profile_image_url,omitempty" toml:"profile_image_url" yaml:"profile_image_url,omitempty"`
	TwitterUserName null.String `boil:"twitter_user_name" json:"twitter_user_name,omitempty" toml:"twitter_user_name" yaml:"twitter_user_name,omitempty"`
	Channel         null.String `boil:"channel" json:"channel,omitempty" toml:"channel" yaml:"channel,omitempty"`
}
