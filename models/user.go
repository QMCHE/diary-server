package models

import (
	"crypto/sha512"
	"encoding/hex"
	"time"

	"gorm.io/gorm"
)

// User is struct of user
type User struct {
	ID        uint            `gorm:"primarykey" json:"id"`
	CreatedAt time.Time       `json:"created_at"`
	UpdatedAt time.Time       `json:"updated_at"`
	DeletedAt *gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
	Name      string          `gorm:"size:45;NOT NULL;" json:"name"`
	UserID    string          `gorm:"size:45;NOT NULL;UNIQUE;" json:"userId"`
	Password  string          `gorm:"size:1000;NOT NULL;" json:"-"`
	IsAdmin   bool            `gorm:"default:false;NOT NULL;" json:"isAdmin"`
}

// IsExists checks is user exists
func (u *User) IsExists(db *gorm.DB) bool {
	encryptedPassword := encryptPassword(u.Password)
	return db.Model(User{}).Where("user_id = ? AND password = ?", u.UserID, encryptedPassword).Take(&u).Error == nil
}

// IsUniqueUserID checks if the userId is unique
func (u *User) IsUniqueUserID(db *gorm.DB) bool {
	return db.Model(User{}).Where("user_id = ?", u.UserID).Take(&u).Error != nil
}

// Create inserts user to db
func (u *User) Create(db *gorm.DB) error {
	u.Password = encryptPassword(u.Password)
	return db.Create(&u).Error
}

// Update updates user
func (u *User) Update(db *gorm.DB) error {
	u.Password = encryptPassword(u.Password)
	return db.Save(&u).Error
}

// Delete deletes user
func (u *User) Delete(db *gorm.DB) error {
	return db.Delete(&u).Error
}

// GetUserByID finds user by id
func (u *User) GetUserByID(db *gorm.DB) error {
	return db.Model(&User{}).Where("id = ?", u.ID).Find(&u).Error
}

// GetUserByUserID finds user by userid and returns error
func (u *User) GetUserByUserID(db *gorm.DB) error {
	return db.Model(&User{}).Where("user_id = ?", u.UserID).Find(&u).Error
}

func encryptPassword(password string) string {
	sha := sha512.New()
	sha.Write([]byte(password))
	return hex.EncodeToString(sha.Sum(nil))
}
