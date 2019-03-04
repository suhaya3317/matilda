package entity

import "time"

type User struct {
	UserID    string    `json:"user_id" datastore:"-" goon:"id"`
	Name      string    `json:"name" datastore:"name"`
	IconPath  string    `json:"icon_path" datastore:"icon_path"`
	Deleted   bool      `json:"deleted" datastore:"deleted"`
	CreatedAt time.Time `json:"created_at" datastore:"created_at"`
	UpdatedAt time.Time `json:"updated_at" datastore:"updated_at"`
}
