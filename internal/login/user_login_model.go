package login

type UserLogin struct {
	Id			 string `db:"id"`
	Email 		 string `db:"email"`
	PasswordHash string `db:"password_hash"`
}