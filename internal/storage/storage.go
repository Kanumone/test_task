package storage

import (
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/kanumone/avito_test/internal/lib/logger"
	"github.com/kanumone/avito_test/internal/storage/models"
)

type Storage struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) *Storage {
	return &Storage{db: db}
}

func (s *Storage) CreateSlug(slug *models.Slug) bool {
	const op = "internal.storage.CreateSlug"
	slug.CreatedAt = time.Now()
	_, err := s.db.NamedExec(`INSERT INTO slugs(title, created_at) VALUES(:title, :created_at) RETURNING id`, slug)
	if err != nil {
		logger.ErrorWrap(op, err.Error())
		return false
	}
	return true
}

func (*Storage) DeleteSlug(title string) {

}

func (*Storage) SlugToUser(username string) {

}

func (*Storage) CreateUser(username string) {

}
