package users

type UserRepository interface {
	GetById(id string) (*User, error)
	Create(user *User) (string, error)
}