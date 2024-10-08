package router

import (
	"github.com/glinboy/fiber-crud-rest-api-demo/model"
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
	books.Post("", func(c *fiber.Ctx) error {
		var book model.Book
		if err := c.BodyParser(&book); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
		}
		bookService.Save(book)
		return c.SendStatus(fiber.StatusCreated)
	})
	books.Put("", func(c *fiber.Ctx) error {
		var book model.Book
		if err := c.BodyParser(&book); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
		}
		bookService.Update(book)
		return c.SendStatus(fiber.StatusOK)
	})
	books.Delete("/:id<int>", func(c *fiber.Ctx) error {
		id, _ := c.ParamsInt("id")
		bookService.Delete(id)
		return c.SendStatus(fiber.StatusNoContent)
	})

	router.Use(func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"msg": "Path doesn't exist",
		})
	})

}
