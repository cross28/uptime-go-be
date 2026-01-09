package login

type UserLogin struct {
	Email 		 string `db:"email"`
	PasswordHash string `db:"password_hash"`
}