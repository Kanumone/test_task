package handlers

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/kanumone/avito_test/internal/lib/api/response"
	"github.com/kanumone/avito_test/internal/lib/helpers"
	"github.com/kanumone/avito_test/internal/lib/logger"
	"github.com/kanumone/avito_test/internal/storage"
	"github.com/kanumone/avito_test/internal/storage/models"
)

type Slugs struct{}

func (sl Slugs) Routes(s *storage.Storage) chi.Router {
	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Get slugs")
	})
	r.Get("/users", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Slug users")
	})
	r.Post("/", func(w http.ResponseWriter, r *http.Request) {
		const op = "server.handlers.slugs.Post"
		slug := models.Slug{}
		err := helpers.ParseJson(r.Body, &slug)
		if err != nil {
			logger.ErrorWrap(op, err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		ok := s.CreateSlug(slug)
		if ok {
			w.Write(response.OK(slug))
		} else {
			w.Write(response.Error("invalid slug"))
		}
	})
	r.Put("/", func(w http.ResponseWriter, r *http.Request) {
		const op = "server.handlers.slugs.Put"
		data := models.UserSlug{}
		err := helpers.ParseJson(r.Body, &data)
		if err != nil {
			logger.ErrorWrap(op, err.Error())
			w.WriteHeader(http.StatusBadRequest)
			response.Error("invalid json")
			return
		}
		ok := s.SlugToUser(data)
		if ok {
			w.Write(response.OK("updated successfully"))
		} else {
			w.Write(response.Error("something went wrong"))
		}
	})
	r.Delete("/", func(w http.ResponseWriter, r *http.Request) {
		const op = "server.handlers.slugs.Delete"
		slug := models.Slug{}
		err := helpers.ParseJson(r.Body, &slug)
		if err != nil {
			logger.ErrorWrap(op, err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		ok := s.DeleteSlug(slug)
		if ok {
			w.Write(response.OK("deleted successfully"))
		} else {
			w.Write(response.Error("invalid slug"))
		}
	})
	return r
}
