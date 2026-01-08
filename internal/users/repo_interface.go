package users

import "context"

type UserRepository interface {
	GetAllUsers(ctx context.Context) ([]User, error)
	GetById(ctx context.Context, id string) (User, error)
	Create(ctx context.Context, user *User) (string, error)
}
