package service

import (
	"github.com/glinboy/fiber-crud-rest-api-demo/model"
	"github.com/glinboy/fiber-crud-rest-api-demo/repository"
)

type BookServiceImpl struct {
	BookRepository repository.BookRepository
}

func NewBookServiceImpl(bookRepository repository.BookRepository) BookService {
	return &BookServiceImpl{
		BookRepository: bookRepository,
	}
}

func (b BookServiceImpl) FindAll() []model.Book {
	return b.BookRepository.FindAll()
}

func (b BookServiceImpl) FindById(id int) model.Book {
	book, err := b.BookRepository.FindById(id)
	if err != nil {
		panic(err)
	}
	return book
}
