package entities

import (
	"time"
)

type Slug struct {
	ID        int64     `db:"id"`
	Title     string    `db:"title"`
	CreatedAt time.Time `db:"created_at"`
}
