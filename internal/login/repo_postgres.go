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

func (repo *PostgresLoginRepo) GetUserByEmail(ctx context.Context, email string) (UserLogin, error) {
	sql := `
		SELECT id, email, password_hash
		FROM public.users
		WHERE email=$1
	`

	rows, err := repo.db.Query(ctx, sql, email)

	if errors.Is(err, pgx.ErrNoRows){
		return UserLogin{}, err
	} else if err != nil {
		return UserLogin{}, fmt.Errorf("failed to retrieve hash: %w", err)
	}

	defer rows.Close()


	user, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[UserLogin])
	if err != nil {
		return UserLogin{}, err
	}

	return user, nil
}
