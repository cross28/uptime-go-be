package users

import "fmt"

type PostgresUserRepository struct {}

func NewPostgresUserRepository() *PostgresUserRepository {
	return &PostgresUserRepository{}
}

func (repo *PostgresUserRepository) GetById(id string) (*User, error) {
	// implement db
	return &User{
		Id: id,
		Email: fmt.Sprintf("%s@crosssystems.co", id),
		PasswordHash: "Th!s1sAP4ssw0rdH4sh!!",
	}, nil
}

func (repo *PostgresUserRepository) Create(user *User) (string, error) {
	// implement db
	return user.Id, nil
}