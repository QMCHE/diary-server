package models

import (
	"crypto/sha512"
	"encoding/hex"

	"gorm.io/gorm"
)

// User is struct of user
type User struct {
	gorm.Model
	Name     string `gorm:"size:45;NOT NULL;" json:"name"`
	UserID   string `gorm:"size:45;NOT NULL;" json:"userId"`
	Password string `gorm:"size:1000;NOT NULL;" json:"password"`
	Diaries  []Diary
}

func encryptPassword(password string) string {
	sha := sha512.New()
	sha.Write([]byte(password))
	return hex.EncodeToString(sha.Sum(nil))
}
