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

func (s *Storage) CreateSlug(slug models.Slug) bool {
	const op = "internal.storage.CreateSlug"
	slug.CreatedAt = time.Now()
	_, err := s.db.NamedExec(`INSERT INTO slugs(title, created_at) VALUES(:title, :created_at) RETURNING id`, slug)
	if err != nil {
		logger.ErrorWrap(op, err.Error())
		return false
	}
	return true
}

func (s *Storage) DeleteSlug(slug models.Slug) bool {
	const op = "internal.storage.DeleteSlug"
	_, err := s.db.NamedExec(`DELETE FROM slugs WHERE title = :title`, slug)
	if err != nil {
		logger.ErrorWrap(op, err.Error())
		return false
	}
	return true
}

func (s *Storage) SlugToUser(data models.UserSlug) bool {
	const op = "internal.storage.SlugToUser"
	tx, err := s.db.Begin()
	if err != nil {
		logger.ErrorWrap(op, err.Error())
		return false
	}
	defer tx.Rollback()
	for _, s := range data.Add {
		_, err := tx.Exec(`INSERT INTO users_slugs(user_id, slug_id) VALUES($1, $2)`, data.UserID, s.ID)
		if err != nil {
			logger.ErrorWrap(op, err.Error())
			return false
		}
	}
	for _, s := range data.Delete {
		_, err := tx.Exec(`DELETE FROM user_slugs WHERE user_id = $1 AND slug_id = $2`, data.UserID, s.ID)
		if err != nil {
			logger.ErrorWrap(op, err.Error())
			return false
		}
	}

	err = tx.Commit()
	if err != nil {
		logger.ErrorWrap(op, err.Error())
		return false
	}
	return true
}
