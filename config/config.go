package config

import (
	"errors"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DatabaseUrl string
}

func Load() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		log.Printf("Warning: %v. Resorting to sysmtem-level environment variables", err)
	}

	dbURL := os.Getenv("TODO_APP_CLI_DB_URL")
	if dbURL == "" {
		log.Fatal("TODO_APP_CLI_DB_URL is not set in environment or .env file.")
		return nil, errors.New("TODO_APP_CLI_DB_URL is required.")
	}

	return &Config{
		DatabaseUrl: dbURL,
	}, nil
}
