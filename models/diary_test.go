package models_test

import (
	"testing"

	"github.com/QMCHE/diary-server/models"
)

func TestInsertDiary(t *testing.T) {
	err := models.InsertDiary("title", "content")
	if err != nil {
		t.Error("Failed to insert diary")
	}
}

func TestUpdateDiary(t *testing.T) {
	err := models.UpdateDiary(1, "title1", "content1")
	if err != nil {
		t.Error("Failed to update diary")
	}
}

func TestDeleteDiary(t *testing.T) {
	err := models.DeleteDiary(1)
	if err != nil {
		t.Error("Failed to delete diary")
	}
}
