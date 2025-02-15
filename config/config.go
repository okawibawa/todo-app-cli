package config

import (
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
		log.Fatal("Error loading environment variables.")
		return nil, err
	}

	return &Config{
		DatabaseUrl: os.Getenv("DB_URL"),
	}, nil
}
