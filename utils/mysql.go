package utils

import (
	"os"

	"github.com/QMCHE/diary-server/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// DBConnect returns db object
func DBConnect() *gorm.DB {
	id := os.Getenv("ID")
	password := os.Getenv("PASSWORD")
	protocol := os.Getenv("PROTOCOL")
	DBAddress := os.Getenv("DB_ADDRESS")
	DBPort := os.Getenv("DB_PORT")
	DBName := os.Getenv("DB_NAME")

	dsn := id + ":" + password + "@" + protocol + "(" + DBAddress + ":" + DBPort + ")/" + DBName + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(&models.User{}, &models.Diary{})
	if err != nil {
		panic(err)
	}

	return db
}
