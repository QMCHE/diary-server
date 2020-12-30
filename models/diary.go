package models

import (
	"gorm.io/gorm"
)

// Diary is struct of diary
type Diary struct {
	gorm.Model
	Title   string `gorm:"size:45;NOT NULL;" json:"title"`
	Content string `gorm:"size:1000;NOT NULL;" json:"content"`
	Author  User
}
