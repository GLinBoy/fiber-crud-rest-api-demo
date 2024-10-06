package router

import (
	"github.com/glinboy/fiber-crud-rest-api-demo/service"
	"github.com/gofiber/fiber/v2"
)

func NewRouter(router fiber.Router, bookService service.BookService) {

	router.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, Fiber!")
	})

	api := router.Group("/api")

	books := api.Group("/books")

	books.Get("", func(c *fiber.Ctx) error {
		books := bookService.FindAll()
		return c.JSON(books)
	})
	books.Get("/:id<int>", func(c *fiber.Ctx) error {
		id, _ := c.ParamsInt("id")
		book := bookService.FindById(id)
		return c.JSON(book)
	})

	router.Use(func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"msg": "Path doesn't exist",
		})
	})

}
