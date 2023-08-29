package storage

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/kanumone/avito_test/internal/lib/helpers"
	"github.com/kanumone/avito_test/internal/server/dto"
	"github.com/kanumone/avito_test/internal/storage/entities"
)

type Storage struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) *Storage {
	return &Storage{db: db}
}

func (s *Storage) CreateSlug(title string) error {
	const op = "internal.storage.CreateSlug"
	res, err := s.db.Exec(`INSERT INTO slugs(title) VALUES($1) ON CONFLICT DO NOTHING`, title)
	if err != nil {
		return helpers.Wrap(op, err)
	}
	num, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if num == 0 {
		return fmt.Errorf("DUPLICATE SLUG")
	}
	return nil
}

func (s *Storage) DeleteSlug(title string) error {
	const op = "internal.storage.DeleteSlug"
	slug := entities.Slug{}
	err := s.db.Get(&slug, `DELETE FROM slugs WHERE title = $1`, title)
	if err != nil {
		return helpers.Wrap(op, err)
	}
	return nil
}

func (s *Storage) SlugToUser(data dto.UserSlug) error {
	const op = "internal.storage.SlugToUser"
	tx, err := s.db.Begin()
	if err != nil {
		return helpers.Wrap(op, err)
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()
	for range data.Add {
		_, err := tx.Exec(`INSERT INTO users_slugs(user_id, slug_id) VALUES($1, $2)`, data.User.ID, 1)
		if err != nil {
			return helpers.Wrap(op, err)
		}
	}
	for range data.Delete {
		_, err := tx.Exec(`DELETE FROM user_slugs WHERE user_id = $1 AND slug_id = $2`, data.User.ID, 1)
		if err != nil {
			return helpers.Wrap(op, err)
		}
	}

	err = tx.Commit()
	if err != nil {
		return helpers.Wrap(op, err)
	}
	return nil
}

func (s *Storage) CreateUser(userID int64) error {
	const op = "internal.storage.CreateUser"
	_, err := s.db.NamedExec(`INSERT INTO users(id) VALUES($1) ON CONFLICT DO NOTHING`, userID)
	if err != nil {
		return helpers.Wrap(op, err)
	}
	return nil
}

func (s *Storage) UserSlugs(userID int64) ([]entities.Slug, error) {
	return []entities.Slug{}, nil
}
