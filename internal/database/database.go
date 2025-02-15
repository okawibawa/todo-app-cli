package database

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/okawibawa/todo-app-cli/config"
)

func InitDb() (*pgxpool.Pool, error) {
	var err error

	cfg, err := config.Load()
	if err != nil {
		return nil, err
	}

	config, err := pgxpool.ParseConfig(cfg.DatabaseUrl)
	if err != nil {
		return nil, err
	}

	config.MaxConns = 10

	dbPool, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		return nil, err
	}

	err = dbPool.Ping(context.Background())
	if err != nil {
		return nil, err
	}

	return dbPool, nil
}

func CloseDb(db *pgxpool.Pool) {
	db.Close()
}
