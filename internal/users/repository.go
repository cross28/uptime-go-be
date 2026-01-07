package users

type IUserRepository interface {
	GetById(id string) (*User, error)
	Create(user *User) (string, error)
}

type UserRepository struct {}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func (repo *UserRepository) GetById(id string) (*User, error) {
	// implement db
	return &User{}, nil
}

func (repo *UserRepository) Create(user *User) (string, error) {
	// implement db
	return "new-id", nil
}