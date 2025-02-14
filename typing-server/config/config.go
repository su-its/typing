package config

import (
	"os"
)

type Config struct {
	Environment string
	DBAddr      string
}

func New() *Config {
	environment := os.Getenv("ENVIRONMENT")
	if environment == "" {
		environment = "production"
	}

	dbAddr := os.Getenv("DB_ADDR")
	if dbAddr == "" {
		dbAddr = "db:3306"
	}

	return &Config{
		Environment: environment,
		DBAddr:      dbAddr,
	}
}
