package server

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/kanumone/avito_test/internal/server/handlers"
	"github.com/kanumone/avito_test/internal/storage"
)

func router(s *storage.Storage) chi.Router {
	r := chi.NewRouter()
	r.Use(middleware.RequestID, middleware.Logger)
	r.Get("/user", handlers.UserSlugs(s))
	r.Put("/user", handlers.UpdateUser(s))
	r.Post("/slug", handlers.CreateSlug(s))
	r.Delete("/slug", handlers.DeleteSlug(s))
	return r
}
