package main

import (
	"github.com/glinboy/fiber-crud-rest-api-demo/router"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	router.NewRouter(app)

	app.Listen(":8080")
}
