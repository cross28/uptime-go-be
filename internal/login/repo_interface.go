package login

import "context"

type LoginRepo interface {
	GetPasswordHash(ctx context.Context, email string) (string, error)
}
