package models

import (
	"crypto/sha512"
	"encoding/hex"

	"gorm.io/gorm"
)

// User is struct of user
type User struct {
	gorm.Model
	Name     string
	UserID   string
	Password string
	Diaries  []Diary
}

func encryptPassword(password string) string {
	sha := sha512.New()
	sha.Write([]byte(password))
	return hex.EncodeToString(sha.Sum(nil))
}
