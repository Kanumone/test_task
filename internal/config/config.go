package config

import (
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/kanumone/avito_test/internal/lib/helpers"
)

type Config struct {
	Env        string `yaml:"env"`
	HttpServer `yaml:"http_server"`
}

type HttpServer struct {
	Host               string        `yaml:"host"`
	Port               string        `yaml:"port"`
	RequestTimeout     time.Duration `yaml:"request_timeout"`
	ConnectrionTimeout time.Duration `yaml:"connection_timeout"`
}

func MustLoad() *Config {
	const op = "config.MustLoad"
	var cfg Config
	path := os.Getenv("CONFIG_PATH")
	if path == "" {
		path = "config/config.yaml"
	}
	err := cleanenv.ReadConfig(path, &cfg)
	if err != nil {
		helpers.LogErr(op, err)
		os.Exit(1)
	}
	return &cfg
}
