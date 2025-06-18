package config

import (
	"fmt"

	"github.com/caarlos0/env/v11"
)

type (
	// Config -.
	Config struct {
		App   App
		Log   Log
		Redis Redis
	}

	// App -.
	App struct {
		Name    string `env:"APP_NAME,required"`
		Version string `env:"APP_VERSION,required"`
	}

	// Redis
	Redis struct {
		Host string `env:"REDIS_HOST,required"`
		Port string `env:"REDIS_PORT,required"`
		// Password string `env:"REDIS_PASSWORD,required"`
		// DB       int    `env:"REDIS_DB,required"`
	}
	// Log -.
	Log struct {
		Level string `env:"LOG_LEVEL,required"`
	}
)

// NewConfig returns app config.
func NewConfig() (*Config, error) {
	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		return nil, fmt.Errorf("config error: %w", err)
	}

	return cfg, nil
}
