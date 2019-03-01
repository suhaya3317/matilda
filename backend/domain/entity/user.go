package entity

import "time"

type User struct {
	UserID    string    `json:"user_id"`
	Name      string    `json:"name"`
	IconPath  string    `json:"icon_path"`
	Deleted   bool      `json:"deleted"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
