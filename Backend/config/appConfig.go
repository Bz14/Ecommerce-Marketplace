package config

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	Port string
}

func LoadConfig() (*AppConfig, error) {
	if os.Getenv("APP_ENV") == "development" {
		godotenv.Load()
	}

	httpPort := os.Getenv("HTTP_PORT")

	if httpPort == "" {
		return &AppConfig{}, errors.New("HTTP_PORT environment variable is not set")
	}

	return &AppConfig{
		Port: httpPort,
	}, nil
}