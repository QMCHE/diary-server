package utils

import (
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// Claims is struct that stores information from tokens
type Claims struct {
	id uint
	jwt.StandardClaims
}

// CreateUserToken returns user's new token
func CreateUserToken(id uint) (string, error) {
	expirationTime := 5 * time.Minute
	key := []byte(os.Getenv("JWT_KEY"))
	claims := &Claims{
		id: id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Milliseconds(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(key)
}
