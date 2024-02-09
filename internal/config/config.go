package config

import (
	"fmt"
	"github.com/caarlos0/env/v10"
	"strconv"
)

type Config struct {
	Postgres struct {
		Host     string `env:"POSTGRES_HOST,notEmpty"`
		Port     string `env:"POSTGRES_PORT,notEmpty"`
		User     string `env:"POSTGRES_USER,notEmpty"`
		Password string `env:"POSTGRES_PASSWORD,notEmpty"`
		Database string `env:"POSTGRES_DB,notEmpty"`
	}
	SMTP struct {
		Host     string `env:"SMTP_HOST"`
		Port     int    `env:"SMTP_PORT"`
		Username string `env:"SMTP_USERNAME"`
		Password string `env:"SMTP_PASSWORD"`
		Sender   string `env:"SMTP_SENDER"`
	}
	Limiter struct {
		Rps     float64 `env:"LIMITER_RPS"`
		Burst   int     `env:"LIMITER_BURST"`
		Enabled bool    `env:"LIMITER_ENABLED"`
	}
	Address string `env:"Address"`
	Port    int    `env:"Port"`
}

func (c *Config) PostgresDSN() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		c.Postgres.Host, c.Postgres.Port, c.Postgres.User, c.Postgres.Password, c.Postgres.Database,
	)
}
func (c *Config) ServerAddress() string {
	return c.Address + ":" + strconv.Itoa(c.Port)
}

func Read() (*Config, error) {
	var config Config

	if err := env.Parse(&config); err != nil {
		return nil, fmt.Errorf("parse config: %w", err)
	}
	return &config, nil
}
