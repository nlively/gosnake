package server

import (
	"fmt"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type ServerConfig struct {
	Host     string `env:"HOST"`
	Port     int    `env:"PORT"`
	MaxGames int    `env:"MAX_GAMES"`
}

func LoadConfig() (*ServerConfig, error) {
	// Load the .env file
	err := godotenv.Load(".env")
	if err != nil {
		return nil, fmt.Errorf("error loading .env file: %w", err)
	}

	var cfg ServerConfig
	err = envconfig.Process("", &cfg)
	if err != nil {
		return nil, fmt.Errorf("error processing env vars: %w", err)
	}

	return &cfg, nil
}
