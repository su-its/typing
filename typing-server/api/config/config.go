package config

import (
	"log/slog"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Environment string
}

func New(logger *slog.Logger) *Config {
	if err := godotenv.Load(); err != nil {
		logger.Warn(".env file not found")
	}

	environment := os.Getenv("ENVIRONMENT")
	if environment == "" {
		environment = "production"
	}

	return &Config{
		Environment: environment,
	}
}
