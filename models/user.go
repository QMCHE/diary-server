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

// IsUserExists checks is user exists
func IsUserExists(db *gorm.DB, userID, password string) error {
	var user User
	encryptedPassword := encryptPassword(password)
	return db.Table("user").Find(&user).Where("userId = ? AND password = ?", userID, encryptedPassword).Error
}

// IsUniqueUserID checks if the userId is unique
func IsUniqueUserID(db *gorm.DB, userID string) bool {
	var id string
	return db.Raw("SELECT id FROM user WHERE userId=?", userID).Row().Scan(&id) != nil
}

// InsertUser inserts user to db
func InsertUser(db *gorm.DB, name, userID, password string) error {
	encryptedPassword := encryptPassword(password)
	user := User{
		Name:     name,
		UserID:   userID,
		Password: encryptedPassword,
		Diaries:  nil,
	}
	err := db.Create(&user).Error
	if err != nil {
		return err
	}
	err = db.Save(&user).Error
	return err
}

func encryptPassword(password string) string {
	sha := sha512.New()
	sha.Write([]byte(password))
	return hex.EncodeToString(sha.Sum(nil))
}
