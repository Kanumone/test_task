package entities

import (
	"time"
)

type Slug struct {
	Title     string    `db:"title"`
	CreatedAt time.Time `db:"created_at"`
}
