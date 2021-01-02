package models_test

import (
	"log"
	"testing"

	"github.com/QMCHE/diary-server/models"
	"github.com/QMCHE/diary-server/utils"
)

func TestIsUserExists(t *testing.T) {
	db := utils.DBConnect()
	user := &models.User{
		Name:     "test",
		Password: "1234",
	}

	err := user.IsUserExists(db)
	if err != nil {
		log.Print(err)
		t.Error(err)
	}
}

func TestIsUniqueUserID(t *testing.T) {
	db := utils.DBConnect()
	user := &models.User{
		UserID: "test",
	}

	if !user.IsUniqueUserID(db) {
		t.Error("'test' is not unique user id")
	}
}

func TestInsertUser(t *testing.T) {
	db := utils.DBConnect()
	user := &models.User{
		Name:     "test",
		UserID:   "test",
		Password: "1234",
	}

	err := user.CreateUser(db)
	if err != nil {
		t.Error(err)
	}
}
