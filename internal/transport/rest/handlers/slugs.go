package handlers

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/kanumone/avito_test/internal/lib/api/response"
	"github.com/kanumone/avito_test/internal/lib/helpers"
	"github.com/kanumone/avito_test/internal/models"
	"github.com/kanumone/avito_test/internal/transport/dto"
)

var validate *validator.Validate = helpers.NewValidator()

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
			response.SendError(w, response.InvalidJson)
			return
		}

		err = validate.Struct(slug)
		if err != nil {
			response.ValidationError(w, err.(validator.ValidationErrors))
			return
		}
		err = c.CreateSlug(slug.Title)
		if err != nil {
			response.SendError(w, err)
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
			response.SendError(w, response.InvalidJson)
			return
		}
		err = d.DeleteSlug(slug.Title)
		if err != nil {
			response.SendError(w, err)
		} else {
			w.Write(response.OK("deleted successfully"))
		}
	}
}
