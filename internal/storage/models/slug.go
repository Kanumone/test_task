package models

import (
	"time"
)

type Slug struct {
	Id        int
	Title     string
	CreatedAt time.Time `db:"created_at"`
}
