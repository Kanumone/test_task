package main

import (
	"github.com/kanumone/avito_test/internal/config"
	"github.com/kanumone/avito_test/internal/storage"
	"github.com/kanumone/avito_test/internal/storage/postgres"
	"github.com/kanumone/avito_test/internal/transport/rest"
)

func main() {
	cfg := config.MustLoad()
	s := storage.New(postgres.New())
	server := rest.Server{}
	server.Start(cfg, s)
}
