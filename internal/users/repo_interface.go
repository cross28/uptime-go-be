package users

import "context"

type UserRepository interface {
	GetById(ctx context.Context, id string) (User, error)
	Create(ctx context.Context, user *User) (string, error)
}
