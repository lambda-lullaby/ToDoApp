package config

import (
	"fmt"
	"time"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	HTTP     HTTPConfig
	Postgres PostgresConfig
}

type HTTPConfig struct {
	Port            int           `envconfig:"HTTP_PORT" default:"8080"`
	ShutdownTimeout time.Duration `envconfig:"HTTP_SHUTDOWN_TIMEOUT" default:"5s"`
}

type PostgresConfig struct {
	Host      string        `envconfig:"POSTGRES_HOST" default:"localhost"`
	Port      int           `envconfig:"POSTGRES_PORT" default:"5432"`
	User      string        `envconfig:"POSTGRES_USER" required:"true"`
	Password  string        `envconfig:"POSTGRES_PASSWORD" required:"true"`
	DB        string        `envconfig:"POSTGRES_DB" required:"true"`
	OpTimeout time.Duration `envconfig:"POSTGRES_OP_TIMEOUT" default:"5s"`
}

func (c PostgresConfig) DSN() string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=disable",
		c.User, c.Password, c.Host, c.Port, c.DB,
	)
}

func NewConfigMust() Config {
	var cfg Config
	if err := envconfig.Process("", &cfg); err != nil {
		panic(fmt.Sprintf("config: %v", err))
	}
	return cfg
}
