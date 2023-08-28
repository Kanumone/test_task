package models

import (
	"time"
)

type Slug struct {
	Id        int64     `db:"id" json:"-"`
	Title     string    `db:"title" json:"title"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
}
