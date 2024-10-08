package repository

import (
	"errors"

	"github.com/glinboy/fiber-crud-rest-api-demo/model"
	"gorm.io/gorm"
)

type BookRepositoryImpl struct {
	DB *gorm.DB
}

func NewBookRepositoryImpl(DB *gorm.DB) BookRepository {
	return &BookRepositoryImpl{DB: DB}
}

func (b BookRepositoryImpl) FindAll() []model.Book {
	var books []model.Book
	result := b.DB.Find(&books)
	if result.Error != nil {
		panic(result.Error)
	}
	return books
}

func (b BookRepositoryImpl) FindById(id int) (model.Book, error) {
	var book model.Book
	result := b.DB.Find(&book, id)
	if result.Error != nil {
		panic(result.Error)
	}
	if result != nil {
		return book, nil
	} else {
		return book, errors.New("Book not found")
	}
}

func (b BookRepositoryImpl) Save(book model.Book) {
	result := b.DB.Create(&book)
	if result.Error != nil {
		panic(result.Error)
	}
}

func (b BookRepositoryImpl) Update(book model.Book) {
	result := b.DB.Save(&book)
	if result.Error != nil {
		panic(result.Error)
	}
}

func (b BookRepositoryImpl) Delete(id int) {
	var book model.Book
	result := b.DB.Where("id = ?", id).Delete(&book)
	if result.Error != nil {
		panic(result.Error)
	}
}
