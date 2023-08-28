package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/kanumone/avito_test/internal/lib/api/response"
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
		const op = "Post.CreateSlug"
		decoder := json.NewDecoder(r.Body)
		// data := make([]byte, 0, 100)
		// r.Body.Read(data)
		// log.Printf("data: %v\n", data)
		slug := models.Slug{}
		err := decoder.Decode(&slug)
		if err != nil {
			logger.ErrorWrap(op, err.Error())
		}
		log.Printf("slug: %v\n", slug)
		ok := s.CreateSlug(&slug)
		if ok {
			w.Write(response.OK(slug))
		}
	})
	r.Put("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Add user to slug")
	})
	return r
}
