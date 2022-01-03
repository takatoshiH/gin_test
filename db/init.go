package db

import (
	. "gin_test/model"
	"github.com/jinzhu/gorm"
)

func Init() {
	db, err := gorm.Open("sqlite3", "gin_test.sqlite")
	if err != nil {
		panic("データベースが開けません（dbInit）")
	}
	db.AutoMigrate(&Album{})
	defer db.Close()
}
