package login

import "context"

type LoginRepo interface {
	GetUserByEmail(ctx context.Context, email string) (UserLogin, error)
}
