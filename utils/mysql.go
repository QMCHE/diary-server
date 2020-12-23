package utils

import (
	"database/sql"
)

// DBConnect returns db object
func DBConnect() *sql.DB {
	db, err := sql.Open("mysql", "root:1234@tcp(127.0.0.1:3306)/diary")

	if err != nil {
		panic(err)
	}

	return db
}
