package router

import "github.com/gofiber/fiber/v2"

func NewRouter(router fiber.Router) {

	router.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, Fiber!")
	})

}
