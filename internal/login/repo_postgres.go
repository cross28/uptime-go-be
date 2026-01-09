package login

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresLoginRepo struct {
	db *pgxpool.Pool
}

func NewPostgresLoginRepo(db *pgxpool.Pool) *PostgresLoginRepo {
	return &PostgresLoginRepo{db: db}
}

func (repo *PostgresLoginRepo) GetPasswordHash(ctx context.Context, email string) (string, error) {
	sql := `
		SELECT password_hash
		FROM public.users
		WHERE email=$1
	`

	var password_hash string
	err := repo.db.QueryRow(ctx, sql, email).Scan(&password_hash)
	if errors.Is(err, pgx.ErrNoRows){
		return "", err
	} else if err != nil {
		return "", fmt.Errorf("failed to retrieve hash: %w", err)
	}

	return password_hash, nil
}
