package utils

import (
	"errors"
	"os"
	"time"

	"github.com/QMCHE/diary-server/models"
	"github.com/dgrijalva/jwt-go"
)

// Claims is struct that stores information from tokens
type Claims struct {
	UserID  string
	IsAdmin bool
	jwt.StandardClaims
}

// GenerateAccessToken returns access token
func GenerateAccessToken(user *models.User) (string, error) {
	expirationTime := 5 * time.Minute
	claims := &Claims{
		UserID:  user.UserID,
		IsAdmin: user.IsAdmin,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Milliseconds(),
		},
	}
	return generateToken(claims)
}

// GenerateRefreshToken returns refresh token
func GenerateRefreshToken(user *models.User) (string, error) {
	expirationTime := 24 * time.Hour * 7
	claims := &Claims{
		IsAdmin: user.IsAdmin,
		StandardClaims: jwt.StandardClaims{
			Issuer:    "localhost",
			ExpiresAt: expirationTime.Milliseconds(),
		},
	}
	return generateToken(claims)
}

// VerifyToken verifies that the token is valid
func VerifyToken(token string) (*jwt.Token, error) {
	claims, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, isValid := t.Method.(*jwt.SigningMethodHMAC); !isValid {
			return nil, errors.New("Unexpected signing method")
		}
		return []byte(os.Getenv("JWT_KEY")), nil
	})
	if err != nil {
		return nil, err
	}
	return claims, nil
}

func generateToken(claims *Claims) (string, error) {
	key := []byte(os.Getenv("JWT_KEY"))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(key)
}
