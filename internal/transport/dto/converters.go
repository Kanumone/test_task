package dto

import (
	"github.com/kanumone/avito_test/internal/storage/entities"
)

func SlugToDTO(e entities.Slug) Slug {
	dto := Slug{
		Title: e.Title,
	}
	return dto
}

func SlugFromDTO(dto Slug) entities.Slug {
	e := entities.Slug{
		Title: dto.Title,
	}
	return e
}

func SlugSliceToDTO(slugs []entities.Slug) []Slug {
	dto := make([]Slug, 0, len([]entities.Slug{}))
	for _, e := range slugs {
		dto = append(dto, SlugToDTO(e))
	}
	return dto
}

func UserToDTO(e entities.User) User {
	dto := User{
		ID: e.ID,
	}
	return dto
}

func UserFromDTO(dto User) entities.User {
	e := entities.User{
		ID: dto.ID,
	}
	return e
}
