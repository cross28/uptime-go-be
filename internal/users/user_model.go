package users

type User struct {
	Id           string `json:"id,omitempty" db:"id"`
	Email        string `json:"email,omitempty" db:"email"`
	PasswordHash string `json:"password_hash,omitempty" db:"passwordhash"`
}
