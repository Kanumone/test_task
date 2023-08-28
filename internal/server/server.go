package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/kanumone/avito_test/internal/config"
	"github.com/kanumone/avito_test/internal/storage"
)

func routes(s *storage.Storage) chi.Router {
	r := chi.NewRouter()
	r.Use(middleware.RequestID, middleware.Logger)

	r.Mount("/slugs", slugs{}.routes(s))
	r.Get("/test", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "TEST!")
	})
	return r
}

func Start(cfg *config.Config, s *storage.Storage) {
	err := http.ListenAndServe(fmt.Sprintf("%s:%s", cfg.Host, cfg.Port), routes(s))
	if err != nil {
		log.Fatal(err)
	}
}
