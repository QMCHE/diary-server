package utils

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// DBConnect returns db object
func DBConnect() *gorm.DB {
	dsn := "root:1234@tcp(127.0.0.1:3306)/diary"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	return db
}
