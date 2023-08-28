package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/kanumone/avito_test/internal/lib/logger"
	"github.com/kanumone/avito_test/internal/storage"
	"github.com/kanumone/avito_test/internal/storage/models"
)

type slugs struct{}

func (self slugs) routes(s *storage.Storage) chi.Router {
	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Get slugs")
	})
	r.Get("/users", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Slug users")
	})
	r.Post("/", func(w http.ResponseWriter, r *http.Request) {
		const op = "Post.CreateSlug"
		body, err := r.GetBody()
		if err != nil {
			logger.ErrorWrap(op, err.Error())
		}
		defer body.Close()
		var data []byte
		body.Read(data)
		var slug models.Slug
		json.Unmarshal(data, slug)
	})
	r.Put("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Add user to slug")
	})
	return r
}
