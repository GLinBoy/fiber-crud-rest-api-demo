package main

import (
	"github.com/glinboy/fiber-crud-rest-api-demo/config"
	"github.com/glinboy/fiber-crud-rest-api-demo/model"
	"github.com/glinboy/fiber-crud-rest-api-demo/repository"
	"github.com/glinboy/fiber-crud-rest-api-demo/router"
	"github.com/glinboy/fiber-crud-rest-api-demo/service"
	"github.com/gofiber/fiber/v2"
)

func main() {

	db := config.DatabaseConnection()

	db.Table("book").AutoMigrate(&model.Book{})

	bookRepository := repository.NewBookRepositoryImpl(db)

	bookService := service.NewBookServiceImpl(bookRepository)

	app := fiber.New()

	router.NewRouter(app, bookService)

	app.Listen(":8080")
}
