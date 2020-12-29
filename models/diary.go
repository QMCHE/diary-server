package models

import (
	"log"

	"github.com/QMCHE/diary-server/utils"
	"gorm.io/gorm"
)

// Diary is struct of diary
type Diary struct {
	gorm.Model
	Title   string
	Content string
	Author  User
}

// GetAllDiaries returns all diaries in DB
func GetAllDiaries() ([]Diary, error) {
	db := utils.DBConnect()
	defer db.Close()

	rows, err := db.Query("SELECT * FROM diary")
	if err != nil {
		log.Print(err)
		return nil, err
	}

	diaries := []Diary{}

	for rows.Next() {
		var id, title, content, author, created, updated string
		if err := rows.Scan(&id, &title, &content, &author, &created, &updated); err != nil {
			return nil, err
		}
		diaries = append(diaries, Diary{id, title, content, author, created, updated})
	}

}

// InsertDiary inserts diary to DB
func InsertDiary(title, content string) error {
	db := utils.DBConnect()
	defer db.Close()

	_, err := db.Exec("INSERT INTO diary (title, content) VALUES (?, ?)", &title, &content)
	return err
}

// UpdateDiary modifies diary
func UpdateDiary(id int, title, content string) error {
	db := utils.DBConnect()
	defer db.Close()

	_, err := db.Exec("UPDATE diary SET title=?, content=? WHERE id=?", &title, &content, &id)
	return err
}

// DeleteDiary deletes diary
func DeleteDiary(id int) error {
	db := utils.DBConnect()
	defer db.Close()

	_, err := db.Exec("DELETE FROM diary WHERE id=?", &id)
	return err
}
