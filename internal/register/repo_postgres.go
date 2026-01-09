package register

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresRegisterRepo struct {
	db *pgxpool.Pool
}

func NewPostgresRegisterRepo(db *pgxpool.Pool) *PostgresRegisterRepo {
	return &PostgresRegisterRepo{db: db}
}

func (repo *PostgresRegisterRepo) Register(ctx context.Context, email string, password_hash string) error {
	sql := `
		INSERT INTO public.users (email, password_hash)
		VALUES ($1, $2)
	`
	_, err := repo.db.Exec(ctx, sql, email, password_hash)

	if err != nil {
		return fmt.Errorf("error inserting into users table: %w", err)
	}

	return nil
}