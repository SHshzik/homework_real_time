package config

import "github.com/spf13/viper"

type (
	// Config -.
	Config struct {
		App     App
		Log     Log
		Redis   Redis
		HTTP    HTTP
		Swagger Swagger
	}

	// App -.
	App struct {
		Name    string
		Version string
	}

	// Log -.
	Log struct {
		Level string
	}

	// Redis -.
	Redis struct {
		Host string
		Port string
	}

	// HTTP -.
	HTTP struct {
		Port string
	}

	Swagger struct {
		Enabled bool
	}
)

// NewConfig returns app config.
func NewConfig(path string) (*Config, error) {
	v := viper.New()
	v.SetConfigFile(path)
	v.AutomaticEnv()

	if err := v.ReadInConfig(); err != nil {
		return nil, err
	}

	var cfg Config
	if err := v.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
