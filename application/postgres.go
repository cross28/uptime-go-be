package application

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresConfig struct {
	Database string
	Username string
	Password string
	Host     string
	Port     string
}

func NewPostgresConnection(ctx context.Context, cfg PostgresConfig) (*pgxpool.Pool, error) {
	connectionString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.Database)

	// db, err := pgx.Connect(ctx, connectionString)
	pool, err := pgxpool.New(ctx, connectionString)
	if err != nil {
		log.Fatalf("failed to connect to postgres: %v", err)
	}
	return pool, nil
}
