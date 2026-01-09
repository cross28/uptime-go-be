package register

import "context"

type RegisterRepo interface {
	Register(ctx context.Context, email string, password_hash string) error
}