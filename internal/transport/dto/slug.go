package dto

type Slug struct {
	Title string `json:"title" validate:"required,notblank"`
}
