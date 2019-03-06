package domain

import (
	"time"

	"google.golang.org/appengine/datastore"
)

type Comment struct {
	CommentID   int64          `json:"comment_id"`
	CommentText string         `json:"comment_text"`
	Mine        bool           `json:"mine"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	UserKey     *datastore.Key `json:"-"`
	UserID      string         `json:"user_id"`
	Name        string         `json:"name"`
	IconPath    string         `json:"icon_path"`
}
