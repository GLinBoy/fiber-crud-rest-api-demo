package config

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func DatabaseConnection() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("database.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database")
	}
	return db
}
