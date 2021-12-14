package main

import (
	"go-db-sqlc/src/database"

	"github.com/gofiber/fiber/v2"
)

func main() {
	database.DBConnect()
	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("hello, world")
	})
	app.Listen(":3000")
}
