package utils_test

import (
	"testing"

	"github.com/QMCHE/diary-server/utils"
)

func TestDBConnect(t *testing.T) {
	db := utils.DBConnect()
	sqlDB, err := db.DB()

	if err != nil {
		t.Error(err)
	}

	err = sqlDB.Ping()
	if err != nil {
		t.Error(err)
	}
}
