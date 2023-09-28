package config

import (
	"fmt"
	"github.com/caarlos0/env/v6"
	"time"
)

type Config struct {
	Postgres struct {
		Host              string `env:"POSTGRES_HOST,notEmpty"`
		Port              string `env:"POSTGRES_PORT,notEmpty"`
		User              string `env:"POSTGRES_USER,notEmpty"`
		Password          string `env:"POSTGRES_PASSWORD,notEmpty"`
		Database          string `env:"POSTGRES_DB,notEmpty"`
		ConnectionTimeout time.Duration
	}
	TelegramToken string
}

func Read() (*Config, error) {
	var config Config

	if err := env.Parse(&config); err != nil {
		return nil, fmt.Errorf("failed to parse config: %w", err)
	}

	return &config, nil
}
