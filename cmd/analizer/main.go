package main

import (
	"github.com/kanumone/avito_test/internal/config"
	"github.com/kanumone/avito_test/internal/server"
	"github.com/kanumone/avito_test/internal/storage"
	"github.com/kanumone/avito_test/internal/storage/postgres"
)

func main() {
	cfg := config.MustLoad()
	s := storage.New(postgres.New())
	server.Start(cfg, s)
}
