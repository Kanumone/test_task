package storage

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/kanumone/avito_test/internal/lib/logger"
)

type Storage struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) *Storage {
	return &Storage{db: db}
}

func (s *Storage) CreateSlug(title string) {
	const op = "internal.storage.CreateSlug"
	res, err := s.db.Exec(`insert into slugs(title) values(?)`, title)
	if err != nil {
		logger.ErrorWrap(op, err.Error())
	}
	fmt.Println(res)
}

func (*Storage) DeleteSlug(title string) {

}

func (*Storage) SlugToUser(username string) {

}

func (*Storage) CreateUser(username string) {

}
