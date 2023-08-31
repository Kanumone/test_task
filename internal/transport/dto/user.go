package dto

type User struct {
	ID int64 `json:"user_id" validate:"required"`
}

type UserSlugReq struct {
	User
	Add    []string `json:"add_slugs"`
	Delete []string `json:"delete_slugs"`
}

type UserSlugRes struct {
	Added   []string `json:"added,omitempty"`
	Deleted []string `json:"deleted,omitempty"`
}
