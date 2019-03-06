package entity

import (
	"time"

	"google.golang.org/appengine/datastore"
)

type Comment struct {
	CommentID   int64          `json:"comment_id" datastore:"-" goon:"id"`
	CommentText string         `json:"comment_text" datastore:"comment_text"`
	Deleted     bool           `json:"deleted" datastore:"deleted"`
	MovieID     int            `json:"movie_id" datastore:"movie_id"`
	UserKey     *datastore.Key `json:"user_key" datastore:"user_key"`
	CreatedAt   time.Time      `json:"created_at" datastore:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at" datastore:"updated_at"`
}
