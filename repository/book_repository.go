package repository

import "github.com/glinboy/fiber-crud-rest-api-demo/model"

type BookRepository interface {
	FindAll() []model.Book
	FindById(id int) (book model.Book, err error)
	Save(book model.Book)
	Update(book model.Book)
	Delete(id int)
}
