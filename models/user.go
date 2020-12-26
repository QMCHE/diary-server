package models

import (
	"database/sql"
	"log"
	"time"

	"github.com/QMCHE/diary-server/utils"
)

// User is struct of user
type User struct {
	ID       int       `db:"id" json:"id" xml:"id"`
	Name     string    `db:"name" json:"name" xml:"name"`
	UserID   string    `db:"userId" json:"userId" xml:"userId"`
	Password string    `db:"password" json:"password" xml:"password"`
	Created  time.Time `db:"created_at" json:"created_at" xml:"create_at"`
	Updated  time.Time `db:"updated_at" json:"updated_at" xml:"update_at"`
	Diaries  []Diary   `db:"diaries" json:"diaries" xml:"diaries"`
}

// IsUserExists checks user exist
func IsUserExists(userID, password string) bool {
	db := utils.DBConnect()
	defer db.Close()

	var user string
	err := db.QueryRow("SELECT userId FROM user WHERE userId=? AND password=?", userID, password).Scan(&user)
	return err == nil
}

// IsUniqueUserID checks userId is unique
func IsUniqueUserID(userID string) bool {
	db := utils.DBConnect()
	defer db.Close()

	return db.QueryRow("SELECT username FROM user").Scan(&userID) != nil
}

// InsertUser inserts user to DB
func InsertUser(name, userID, password string) error {
	db := utils.DBConnect()
	defer db.Close()

	_, err := db.Exec("INSERT INTO user VALUES (?, ?, ?)", name, userID, password)
	return err
}

// GetUserInfoByID returns user info by id
func GetUserInfoByID(id int) *sql.Rows {
	db := utils.DBConnect()
	defer db.Close()

	user, _ := db.Query("SELECT * FROM user WHERE id=?", id)
	defer user.Close()

	var u User

	for user.Next() {
		err := user.Scan(&u.ID, &u.Name, &u.UserID, &u.Password, &u.Created, &u.Updated, &u.Diaries)
		if err != nil {
			log.Print(err)
			return nil
		}
		log.Print(u)
	}
	return user
}

// GetUserInfoByUsername returns user info by username
func GetUserInfoByUsername(username string) *sql.Row {
	db := utils.DBConnect()
	defer db.Close()

	return db.QueryRow("SELECT * FROM user WHERE username=?", username)
}
