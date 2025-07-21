package config

import (
	"fmt"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	LocalPort   int    `envconfig:"LOCAL_PORT" required:"true"`
	PeerAddress string `envconfig:"PEER_ADDRESS" required:"true"`
	PeerPort    int    `envconfig:"PEER_PORT" required:"true"`
}

func LoadConfig() (*Config, error) {
	// Load the .env file
	err := godotenv.Load(".env")
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
