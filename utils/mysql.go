package utils

import (
	"os"

	"github.com/QMCHE/diary-server/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// DBConnect returns db object
func DBConnect() *gorm.DB {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	id := os.Getenv("ID")
	password := os.Getenv("PASSWORD")
	protocol := os.Getenv("PROTOCOL")
	DBAddress := os.Getenv("DBADDRESS")
	DBPort := os.Getenv("DBPORT")
	DBName := os.Getenv("DBNAME")

	dsn := id + ":" + password + "@" + protocol + "(" + DBAddress + ":" + DBPort + ")/" + DBName
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&models.User{}, &models.Diary{})

	return db
}
