package repository

import (
	"gorm.io/gorm"
)

type BookRepositoryImpl struct {
	DB *gorm.DB
}

func NewBookRepositoryImpl(DB *gorm.DB) BookRepository {
	return &BookRepositoryImpl{DB: DB}
}
