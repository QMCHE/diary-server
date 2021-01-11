package utils

import (
	"log"
	"os"

	"github.com/QMCHE/diary-server/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// DBConnect returns db object
func DBConnect() (*gorm.DB, error) {
	wd, err := os.Getwd()
	if err != nil {
		log.Print(err)
	}

	err = godotenv.Load(wd + "/.env")
	if err != nil {
		log.Print(err)
	}

	id := os.Getenv("ID")
	password := os.Getenv("PASSWORD")
	protocol := os.Getenv("PROTOCOL")
	DBAddress := os.Getenv("DB_ADDRESS")
	DBPort := os.Getenv("DB_PORT")
	DBName := os.Getenv("DB_NAME")

	dsn := id + ":" + password + "@" + protocol + "(" + DBAddress + ":" + DBPort + ")/" + DBName + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Print(err)
		return nil, err
	}

	err = db.AutoMigrate(&models.User{}, &models.Diary{})
	if err != nil {
		log.Print(err)
		return nil, err
	}

	return db, nil
}
