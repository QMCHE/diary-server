package models

import (
	"time"

	"github.com/QMCHE/diary-server/utils"
)

// Diary is struct of diary
type Diary struct {
	ID      int       `db:"id" json:"id" xml:"id"`
	Title   string    `db:"title" json:"title" xml:"title"`
	Content string    `db:"content" json:"content" xml:"content"`
	Author  string    `db:"author" json:"author" xml:"author"`
	Created time.Time `db:"created_at" json:"created_at" xml:"created_at"`
	Updated time.Time `db:"updated_at" json:"updated_at" xml:"updated_at"`
}

// InsertDiary inserts diary to DB
func InsertDiary(title, content string) error {
	db := utils.DBConnect()
	defer db.Close()

	_, err := db.Exec("INSERT INTO diary VALUES (?, ?)", &title, &content)
	return err
}

// UpdateDiary modifies diary
func UpdateDiary(id int, title, content string) error {
	db := utils.DBConnect()
	defer db.Close()

	_, err := db.Exec("UPDATE diary SET title=? content=? WHERE id=?", &title, &content, &id)
	return err
}

// DeleteDiary deletes diary
func DeleteDiary(id int) error {
	db := utils.DBConnect()
	defer db.Close()

	_, err := db.Exec("DELETE FROM diary WHERE id=?", &id)
	return err
}
