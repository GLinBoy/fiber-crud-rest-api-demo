package config

import (
	"github.com/jinzhu/gorm"
	"github.com/jinzhu/gorm/dialects/sqlite"
	"gorm.io/driver/sqlite"
)

func DatabaseConnection() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("database.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database")
	}
	return db
}
