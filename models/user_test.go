package models_test

import (
	"testing"

	"github.com/QMCHE/diary-server/models"
)

func TestIsUserExists(t *testing.T) {
	user := models.IsUserExists("test", "1234")
	if user {
		t.Error("User not exist")
	}
}

func TestIsUniqueUserID(t *testing.T) {
	isUnique := models.IsUniqueUserID("test")
	if !isUnique {
		t.Error("UserID is not unique")
	}
}

func TestInsertUser(t *testing.T) {
	result := models.InsertUser("test", "test1", "1234")
	if result != nil {
		t.Error("Failed to insert user")
	}
}
