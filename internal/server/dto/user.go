package dto

type User struct {
	ID int64 `json:"user_id"`
}

type UserSlug struct {
	User
	Add    []string `json:"add_slugs"`
	Delete []string `json:"delete_slugs"`
}
