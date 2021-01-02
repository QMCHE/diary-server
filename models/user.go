package models

import (
	"crypto/sha512"
	"encoding/hex"

	"gorm.io/gorm"
)

// User is struct of user
type User struct {
	gorm.Model
	Name     string  `gorm:"size:45;NOT NULL;" json:"name"`
	UserID   string  `gorm:"size:45;NOT NULL;UNIQUE;" json:"userId"`
	Password string  `gorm:"size:1000;NOT NULL;" json:"password"`
	Diaries  []Diary `gorm:"foreignKey:ID;references:ID"`
}

// BeforeCreate is a function that performs password encryption before the user is saved
func (u *User) BeforeCreate() {
	encryptedPassword := encryptPassword(u.Password)
	u.Password = encryptedPassword
}

// IsUserExists checks is user exists
func (u *User) IsUserExists(db *gorm.DB) error {
	encryptedPassword := encryptPassword(u.Password)
	return db.Model(User{}).Where("user_id = ? AND password = ?", u.UserID, encryptedPassword).Take(&u).Error
}

// IsUniqueUserID checks if the userId is unique
func (u *User) IsUniqueUserID(db *gorm.DB) bool {
	return db.Model(User{}).Where("user_id = ?", u.UserID).Take(&u).Error != nil
}

// CreateUser inserts user to db
func (u *User) CreateUser(db *gorm.DB) error {
	return db.Create(&u).Error
}

func encryptPassword(password string) string {
	sha := sha512.New()
	sha.Write([]byte(password))
	return hex.EncodeToString(sha.Sum(nil))
}
