package auth

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JwtHeader struct {
	Alg string `json:"alg"`
	Typ string `json:"typ"`
}

type UserClaims struct {
	jwt.RegisteredClaims
	UserId string `json:"uid"`
}


func CreateJwtToken(user_id string) (string, error) {
	var secretKey = []byte(os.Getenv("JWT_SECRET_KEY"))

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, 
		jwt.MapClaims{
			"sub": user_id,
			"exp": time.Now().Add(time.Hour * 1).Unix(),
		})

	tokenString, err := token.SignedString(secretKey)

	if err != nil {
		return "", fmt.Errorf("error creating jwt: %v", err)
	}

	return tokenString, nil
}

func VerifyJwtToken(tokenString string) error {
	var secretKey = []byte(os.Getenv("JWT_SECRET_KEY"))

	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (any, error) {
		return secretKey, nil
	})

	if err != nil {
		return fmt.Errorf("error validating jwt: %v", err)
	}

	if !token.Valid {
		return fmt.Errorf("invalid token")
	}
	
	return nil
}