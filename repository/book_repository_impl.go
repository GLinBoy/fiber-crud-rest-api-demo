package repository

import (
	"github.com/glinboy/fiber-crud-rest-api-demo/model"
	"gorm.io/gorm"
)

type BookRepositoryImpl struct {
	DB *gorm.DB
}

func NewBookRepositoryImpl(DB *gorm.DB) BookRepository {
	return &BookRepositoryImpl{DB: DB}
}

func (b BookRepositoryImpl) Save(book model.Book) {
	result := b.DB.Create(&book)
	if result.Error != nil {
		panic(result.Error)
	}
}
