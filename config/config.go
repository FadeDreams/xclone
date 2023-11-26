package config

import (
	"os"

	"github.com/joho/godotenv"
)

type database struct {
	URL string
}

type Config struct {
	Database database
}

func New() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	return &Config{
		Database: database{
			URL: os.Getenv("DATABASE_URL"),
		},
	}, nil
}
