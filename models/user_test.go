package models_test

import (
	"log"
	"testing"

	"github.com/QMCHE/diary-server/models"
	"github.com/QMCHE/diary-server/utils"
)

func TestIsUserExists(t *testing.T) {
	db := utils.DBConnect()

	err := models.IsUserExists(db, "test", "1234")
	if err != nil {
		log.Print(err)
		t.Error(err)
	}
}

func TestIsUniqueUserID(t *testing.T) {
	db := utils.DBConnect()

	if !models.IsUniqueUserID(db, "test") {
		t.Error("'test' is not unique user id")
	}
}

func TestInsertUser(t *testing.T) {
	db := utils.DBConnect()

	err := models.InsertUser(db, "test", "test", "1234")
	if err != nil {
		t.Error(err)
	}
}
