package handlers

import (
	"log"
	"net/http"

	"github.com/kanumone/avito_test/internal/lib/api/response"
	"github.com/kanumone/avito_test/internal/lib/helpers"
	"github.com/kanumone/avito_test/internal/models"
	"github.com/kanumone/avito_test/internal/server/dto"
)

type Creator interface {
	CreateSlug(title string) error
}

func CreateSlug(c Creator) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "server.handlers.slugs.Create"
		slug := dto.Slug{}
		err := helpers.ParseJson(r.Body, &slug)
		if err != nil {
			helpers.LogErr(op, err)
			w.WriteHeader(http.StatusBadRequest)
			w.Write(response.Error("invalid json"))
			return
		}
		err = c.CreateSlug(slug.Title)
		log.Print(err)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write(response.Error(err.Error()))
		} else {
			w.Write(response.OK("created successfully"))
		}
	}
}

type Deletor interface {
	DeleteSlug(title string) error
}

func DeleteSlug(d Deletor) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "server.handlers.slugs.Delete"
		slug := models.Slug{}
		err := helpers.ParseJson(r.Body, &slug)
		if err != nil {
			helpers.LogErr(op, err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		err = d.DeleteSlug(slug.Title)
		if err != nil {
			w.Write(response.OK("deleted successfully"))
		} else {
			w.Write(response.Error("something went wrong"))
		}
	}
}
