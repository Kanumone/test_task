package postgres

import (
	"fmt"
	"log"

	"github.com/ilyakaznacheev/cleanenv"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"github.com/kanumone/avito_test/internal/lib/logger"
)

type Config struct {
	Host     string `env:"DB_HOST" env-default:"localhost"`
	Port     string `env:"DB_PORT"`
	Username string `env:"DB_USER"`
	Password string `env:"DB_PASSWORD"`
	DbName   string `env:"DB_NAME"`
}

func url(c *Config) string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		c.Username, c.Password, c.Host, c.Port, c.DbName)
}

func New() *sqlx.DB {
	const op = "storage.postgres.New"
	var cfg Config
	godotenv.Load()
	cleanenv.ReadEnv(&cfg)
	fmt.Println(cfg)
	db, err := sqlx.Connect("pgx", url(&cfg))
	if err != nil {
		log.Fatal(logger.ErrorWrap(op, err.Error()))
	}
	return db
}
