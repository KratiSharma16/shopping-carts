package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3" // SQLite driver
)

var DB *gorm.DB // global DB connection

func Connect() {
	db, err := gorm.Open("sqlite3", "shopping.db")
	if err != nil {
		panic("failed to connect to database")
	}
	DB = db
}
