package users

import (
	"context"
	"crypto/rand"
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
		FROM public.Users
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
		INSERT INTO public.Users
		(
			Id,
			Email,
			PasswordHash
		)
		VALUES
		(
			$1,
			$2,
			$3
		)
		RETURNING Id
	`

	var id string
	randomId := rand.Text()[:16]
	user.Id = randomId

	err := repo.db.QueryRow(ctx, sql, user.Id, user.Email, user.PasswordHash).Scan(&id)

	if err != nil {
		return "", fmt.Errorf("error inserting user: %w", err)
	}

	return id, nil
}

func (repo *PostgresUserRepository) GetAllUsers(ctx context.Context) ([]User, error) {
	sql := `
		SELECT Id, Email, PasswordHash
		FROM public.Users
	`

	rows, err := repo.db.Query(ctx, sql)
	if err != nil {
		return []User{}, fmt.Errorf("error getting all users: %w", err)
	}
	defer rows.Close()

	users, err := pgx.CollectRows(rows, pgx.RowToStructByName[User])
	if err != nil {
		return []User{}, fmt.Errorf("error collecting rows for get all users: %w", err)
	}

	return users, nil
}
