package application

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresConfig struct {
	ConnectionString string
}

func NewPostgresConnection(ctx context.Context, cfg PostgresConfig) (*pgxpool.Pool, error) {
	pool, err := pgxpool.New(ctx, cfg.ConnectionString)
	if err != nil {
		log.Fatalf("failed to connect to postgres: %v", err)
	}
	return pool, nil
}
