package config

import (
	"fmt"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	PeerAddress string `env:PEER_ADDRESS`
	PeerPort    int    `env:PEER_PORT`
}

func LoadConfig() (*Config, error) {
	// Load the .env fiole
	err := godotenv.Load()
	if err != nil {
		return nil, fmt.Errorf("error loading .env file: %w", err)
	}

	var cfg Config
	err = envconfig.Process("", &cfg)
	if err != nil {
		return nil, fmt.Errorf("error processing env vars: %w", err)
	}

	return &cfg, nil
}
