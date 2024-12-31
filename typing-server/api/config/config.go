package config

import (
	"log/slog"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Environment string
	DBAddr      string
}

func New(logger *slog.Logger) *Config {
	if err := godotenv.Load(); err != nil {
		logger.Warn(".env file not found")
	}

	environment := os.Getenv("ENVIRONMENT")
	if environment == "" {
		environment = "production"
	}

	dbAddr := os.Getenv("DB_ADDR")
	if dbAddr == "" {
		dbAddr = "db:3306"
	}

	logger.Info("config",
		"environment", environment,
		"db_addr", dbAddr)

	return &Config{
		Environment: environment,
		DBAddr:      dbAddr,
	}
}
