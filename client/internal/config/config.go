package config

import (
	"fmt"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type ClientConfig struct {
	ServerHost string `env:"SERVER_HOST"`
	ServerPort int    `env:"SERVER_PORT"`
}

func LoadConfig() (*ClientConfig, error) {
	// Load the .env file
	err := godotenv.Load(".env")
	if err != nil {
		return nil, fmt.Errorf("error loading .env file: %w", err)
	}

	var cfg ClientConfig
	err = envconfig.Process("", &cfg)
	if err != nil {
		return nil, fmt.Errorf("error processing env vars: %w", err)
	}

	return &cfg, nil
}
