package config

import (
	"log"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
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
	err := cleanenv.ReadConfig(path, &cfg)
	if err != nil {
		log.Fatal(op, err)
	}
	return &cfg
}
