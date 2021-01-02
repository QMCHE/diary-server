package models_test

import (
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

	if !user.IsUserExists(db) {
		t.Error("User not exists")
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
