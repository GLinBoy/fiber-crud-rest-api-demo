package router

import "github.com/gofiber/fiber/v2"

func NewRouter(router fiber.Router) {

	router.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, Fiber!")
	})

	router.Use(func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"msg": "Path doesn't exist",
		})
	})

}
