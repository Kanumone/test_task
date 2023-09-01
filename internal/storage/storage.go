package storage

import (
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/kanumone/avito_test/internal/lib/api/response"
	"github.com/kanumone/avito_test/internal/lib/helpers"
	"github.com/kanumone/avito_test/internal/storage/entities"
	"github.com/kanumone/avito_test/internal/transport/dto"
)

type Storage struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) *Storage {
	return &Storage{db: db}
}

// CreateSlug creates a slug for the given title.
//
// title: the title to create a slug for.
// error: an error if there was a problem creating the slug.
func (s *Storage) CreateSlug(title string) error {
	const op = "internal.storage.CreateSlug"
	res, err := s.db.Exec(`INSERT INTO slugs (title) VALUES ($1) ON CONFLICT (title) DO UPDATE SET deleted = false`, title)
	if err != nil {
		helpers.LogErr(op, err)
		return err
	}
	num, err := res.RowsAffected()
	if err != nil {
		helpers.LogErr(op, err)
		return err
	}

	if num == 0 {
		return response.DuplicateErr
	}
	return nil
}

func (s *Storage) DeleteSlug(title string) error {
	const op = "internal.storage.DeleteSlug"
	res, err := s.db.Exec(`UPDATE slugs SET deleted = true, deleted_at = now() WHERE title = $1`, title)
	if err != nil {
		helpers.LogErr(op, err)
		return err
	}
	if ok, _ := res.RowsAffected(); ok == 0 {
		return response.NotFoundErr
	}
	return nil
}

func checkSlugs(tx *sqlx.Tx, slugs []string) (int, error) {
	const op = "internal.storage.checkSlugs"
	query := `SELECT title FROM slugs WHERE title = ANY($1)`
	selected := []string{}
	err := tx.Select(&selected, query, slugs)
	l := len(selected)
	if err != nil {
		helpers.LogErr(op, err)
		return 0, err
	}
	return l, nil
}

func addUserSlugs(tx *sqlx.Tx, data dto.UserSlugReq) ([]string, error) {
	const op = "internal.storage.AddUserSlugs"
	exists, err := checkSlugs(tx, data.Add)
	if len(data.Add) != exists {
		return nil, response.NotAddedSlugs
	}
	if err != nil {
		helpers.LogErr(op, err)
		return nil, err
	}
	query := make([]string, 0, len(data.Add))
	addSlice := make([]interface{}, 0, len(data.Add))
	for i, slug := range data.Add {
		query = append(query, fmt.Sprintf("(%d, $%d)", data.ID, i+1))
		addSlice = append(addSlice, slug)
	}
	result := make([]string, 0, len(data.Add))
	q := fmt.Sprintf("INSERT INTO users_slugs(user_id, slug) VALUES %s ON CONFLICT DO NOTHING RETURNING slug", strings.Join(query, ","))
	err = tx.Select(&result, q, addSlice...)
	if err != nil {
		helpers.LogErr(op, err)
	}
	return result, err
}

func deleteUserSlugs(tx *sqlx.Tx, data dto.UserSlugReq) ([]string, error) {
	const op = "internal.storage.DeleteUserSlugs"
	exists, err := checkSlugs(tx, data.Delete)
	if len(data.Delete) != exists {
		return nil, response.NotDeletedSlugs
	}
	if err != nil {
		helpers.LogErr(op, err)
		return nil, err
	}
	result := make([]string, 0, len(data.Delete))
	query := "DELETE FROM users_slugs WHERE user_id = $1 AND slug = ANY($2) RETURNING slug"
	err = tx.Select(&result, query, data.ID, data.Delete)
	if err != nil {
		helpers.LogErr(op, err)
	}
	return result, err
}

func (s *Storage) SlugToUser(data dto.UserSlugReq) (res entities.AddedDeleted, err error) {
	const op = "internal.storage.SlugToUser"
	err = s.CreateUser(data.User.ID)
	if err != nil {
		return
	}
	res = entities.AddedDeleted{}
	tx, err := s.db.Beginx()
	defer func() {
		if err != nil {
			helpers.LogErr(op, err)
			tx.Rollback()
		}
	}()
	if err != nil {
		return
	}
	if len(data.Add) > 0 {
		res.Added, err = addUserSlugs(tx, data)
		if err != nil {
			return
		}
	}
	if len(data.Delete) > 0 {
		res.Deleted, err = deleteUserSlugs(tx, data)
		if err != nil {
			return
		}
	}

	err = tx.Commit()
	if err != nil {
		helpers.LogErr(op, err)
	}
	return
}

func (s *Storage) CreateUser(userID int64) error {
	const op = "internal.storage.CreateUser"
	_, err := s.db.Exec(`INSERT INTO users(id) VALUES($1) ON CONFLICT DO NOTHING`, userID)
	if err != nil {
		helpers.LogErr(op, err)
		return err
	}
	return nil
}

func (s *Storage) UserSlugs(userID int64) ([]entities.Slug, error) {
	const op = "internal.storage.UserSlugs"
	slugs := []entities.Slug{}
	query := `SELECT slug as title FROM users_slugs
		JOIN slugs ON users_slugs.slug = slugs.title
		WHERE user_id = $1
		AND deleted = false`
	err := s.db.Select(&slugs, query, userID)
	if err != nil {
		helpers.LogErr(op, err)
		return nil, err
	}
	return slugs, nil
}
