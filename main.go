package main

import (
	"go-db-sqlc/src/database"
	"go-db-sqlc/src/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	database.DBConnect()

	app := fiber.New()
	routes.Setup(app)
	app.Listen(":3000")

}
