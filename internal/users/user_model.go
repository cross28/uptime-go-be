package users

type User struct {
	Id string `json:"id"`
	Email string `json:"email"`
	PasswordHash string `json:"password_hash"`
}