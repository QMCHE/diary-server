package models

import (
	"time"

	"gorm.io/gorm"
)

// Diary is struct of diary
type Diary struct {
	ID        uint            `gorm:"primarykey" json:"id"`
	CreatedAt time.Time       `json:"created_at"`
	UpdatedAt time.Time       `json:"updated_at"`
	DeletedAt *gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
	Title     string          `gorm:"size:45;NOT NULL;" json:"title"`
	Content   string          `gorm:"size:1000;NOT NULL;" json:"content"`
	UserID    uint            `json:"userId"`
	User      User            `gorm:"foreignKey:UserID" json:"user"`
}

// GetDiary returns diaries
func GetDiary(db *gorm.DB, sort, direction string, perPage, page int) ([]Diary, error) {
	var diaries []Diary
	err := db.Model(&Diary{}).Order(sort + " " + direction).Limit(perPage).Offset(perPage * (page - 1)).Preload("User").Find(&diaries).Error
	if err != nil {
		return nil, err
	}
	return diaries, nil
}

// Create inserts diary
func (d *Diary) Create(db *gorm.DB) error {
	return db.Create(&d).Error
}

// Update updates diary
func (d *Diary) Update(db *gorm.DB) error {
	return db.Save(&d).Error
}

// Delete deletes diary
func (d *Diary) Delete(db *gorm.DB) error {
	return db.Delete(&d).Error
}

// GetDiaryByID finds diary by id and returns error
func (d *Diary) GetDiaryByID(db *gorm.DB) error {
	return db.Model(&Diary{}).Where("id = ?", d.ID).Find(&d).Error
}
