package config

import (
	"github.com/codingconcepts/env"
)

type Config struct {
	Address      string `env:"ADDRESS" default:"0.0.0.0"`
	Port         string `env:"PORT" default:"5555"`
	DataPath     string `env:"DATA_PATH" default:"/etc/data"`
	LoginAddress string `env:"LOGIN_ADDRESS" default:"localhost:50051"`
	StatePath    string `env:"STATE_PATH" default:"state"`
}

func Load() (*Config, error) {
	config := &Config{}
	if err := env.Set(config); err != nil {
		return nil, err
	}

	return config, nil
}
