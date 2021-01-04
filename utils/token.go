package utils

import (
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// Claims is struct that stores information from tokens
type Claims struct {
	id     uint
	UserID string
	jwt.StandardClaims
}

// GenerateAccessToken returns access token
func GenerateAccessToken(id uint, userID string) (string, error) {
	expirationTime := 5 * time.Minute
	claims := &Claims{
		id:     id,
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Milliseconds(),
		},
	}
	return generateToken(claims)
}

// GenerateRefreshToken returns refresh token
func GenerateRefreshToken(id uint, userID string) (string, error) {
	expirationTime := 24 * time.Hour * 7
	claims := &Claims{
		id:     id,
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Milliseconds(),
		},
	}
	return generateToken(claims)
}

func generateToken(claims *Claims) (string, error) {
	key := []byte(os.Getenv("JWT_KEY"))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(key)
}
