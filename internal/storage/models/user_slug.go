package models

type UserSlug struct {
	UserID int64  `json:"user_id"`
	Add    []Slug `json:"add_slugs"`
	Delete []Slug `json:"delete_slugs"`
}
