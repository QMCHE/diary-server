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
	expirationTime := time.Now().Add(5 * time.Minute).Unix()
	claims := &Claims{
		UserID:  user.UserID,
		IsAdmin: user.IsAdmin,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime,
		},
	}
	return generateToken(claims)
}

// GenerateRefreshToken returns refresh token
func GenerateRefreshToken(user *models.User) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour * 7).Unix()
	claims := &Claims{
		IsAdmin: user.IsAdmin,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime,
		},
	}
	return generateToken(claims)
}

// VerifyToken verifies that the token is valid
func VerifyToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(t *jwt.Token) (interface{}, error) {
		if _, isValid := t.Method.(*jwt.SigningMethodHMAC); !isValid {
			return nil, errors.New("Unexpected signing method")
		}
		return []byte(os.Getenv("JWT_KEY")), nil
	})

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, err
}

func generateToken(claims *Claims) (string, error) {
	key := []byte(os.Getenv("JWT_KEY"))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(key)

}
