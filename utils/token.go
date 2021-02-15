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
	ID      uint
	IsAdmin bool
	jwt.StandardClaims
}

// GenerateAccessToken returns access token
func GenerateAccessToken(user *models.User) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour).Unix()
	claims := &Claims{
		ID:      user.ID,
		IsAdmin: user.IsAdmin,
		StandardClaims: jwt.StandardClaims{
			Issuer:    "localhost",
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: expirationTime,
		},
	}
	return generateToken(claims)
}

// GenerateRefreshToken returns refresh token
// func GenerateRefreshToken(user *models.User) (string, error) {
// 	expirationTime := time.Now().Add(24 * time.Hour * 7).Unix()
// 	claims := &Claims{
// 		ID:      user.ID,
// 		IsAdmin: user.IsAdmin,
// 		StandardClaims: jwt.StandardClaims{
// 			Issuer:    "localhost",
// 			IssuedAt:  time.Now().Unix(),
// 			ExpiresAt: expirationTime,
// 		},
// 	}
// 	return generateToken(claims)
// }

// VerifyToken verifies that the token is valid
func VerifyToken(tokenString string) (*Claims, error) {
	var claims jwt.Claims = &Claims{}
	_, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
		if _, isValid := t.Method.(*jwt.SigningMethodHMAC); !isValid {
			return nil, errors.New("Unexpected signing method")
		}
		return []byte(os.Getenv("JWT_KEY")), nil
	})

	return claims.(*Claims), err
}

// IsExpired checks the token has expired
func IsExpired(tokenString string) bool {
	claims, err := VerifyToken(tokenString)

	v, _ := err.(*jwt.ValidationError)

	if v.Errors == jwt.ValidationErrorClaimsInvalid && claims.ExpiresAt > time.Now().Unix() {
		return false
	}
	return true
}

func generateToken(claims *Claims) (string, error) {
	key := []byte(os.Getenv("JWT_KEY"))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(key)
}
