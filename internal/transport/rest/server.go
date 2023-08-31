package rest

import (
	"fmt"
	"log"
	"net/http"

	"github.com/kanumone/avito_test/internal/config"
	"github.com/kanumone/avito_test/internal/storage"
)

type Server struct{}

func (srv Server) Start(cfg *config.Config, s *storage.Storage) {
	err := http.ListenAndServe(fmt.Sprintf("%s:%s", cfg.Host, cfg.Port), router(s))
	if err != nil {
		log.Fatal(err)
	}
}
