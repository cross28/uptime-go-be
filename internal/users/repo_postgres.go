package users

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresUserRepository struct {
	db *pgxpool.Pool
}

func NewPostgresUserRepository(db *pgxpool.Pool) *PostgresUserRepository {
	return &PostgresUserRepository{db: db}
}

func (repo *PostgresUserRepository) GetById(ctx context.Context, id string) (User, error) {
	sql := `
		SELECT Id, Email, PasswordHash
		FROM dbo.Users
		WHERE Id=$1
	`
	rows, err := repo.db.Query(ctx, sql, id)
	if err != nil {
		return User{}, fmt.Errorf("failed to get user by id: %v", err)
	}
	defer rows.Close()

	user, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[User])

	if err != nil {
		return User{}, fmt.Errorf("error getting user by id: %w", err)
	}

	return user, nil
}

func (repo *PostgresUserRepository) Create(ctx context.Context, user *User) (string, error) {
	sql := `
		INSERT INTO dbo.Users
		(
			Email,
			PasswordHash
		)
		VALUES
		(
			$1,
			$2
		)
		RETURNING Id
	`
	var id string
	err := repo.db.QueryRow(ctx, sql, user.Email, user.PasswordHash).Scan(&id)

	if err != nil {
		return "", fmt.Errorf("error inserting user: %w", err)
	}

	return id, nil
}
