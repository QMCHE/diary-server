package models_test

import (
	"testing"

	"github.com/QMCHE/diary-server/models"
	"github.com/QMCHE/diary-server/utils"
)

func TestIsUserExists(t *testing.T) {
	db, err := utils.DBConnect()
	if err != nil {
		t.Error(err)
	}

	user := &models.User{
		Name:     "test",
		Password: "1234",
	}

	if !user.IsExists(db) {
		t.Error("User not exists")
	}
}

func TestIsUniqueUserID(t *testing.T) {
	db, err := utils.DBConnect()
	if err != nil {
		t.Error(err)
	}

	user := &models.User{
		UserID: "test",
	}

	if !user.IsUniqueUserID(db) {
		t.Error("'test' is not unique user id")
	}
}

func TestInsertUser(t *testing.T) {
	db, err := utils.DBConnect()
	if err != nil {
		t.Error(err)
	}

	user := &models.User{
		Name:     "test",
		UserID:   "test",
		Password: "1234",
	}

	err = user.Create(db)
	if err != nil {
		t.Error(err)
	}
}
