package service

import "github.com/glinboy/fiber-crud-rest-api-demo/model"

type BookService interface {
	FindAll() []model.Book
	FindById(id int) model.Book
	Save(book model.Book)
	Update(book model.Book)
	Delete(id int)
}
