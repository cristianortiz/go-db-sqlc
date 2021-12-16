package main

import (
	"context"
	"database/sql"
	"go-db-sqlc/src/database"
	"go-db-sqlc/src/database/users"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	ctx := context.Background()
	db := database.DBConnect()
	queries := users.New(db)

	// create an user
	result, err := queries.CreateUser(ctx, users.CreateUserParams{
		Firstname:    "Jon",
		Lastname:     "Doe",
		Email:        "jon@gmail.com",
		Upassword:    "1234",
		Isambassador: sql.NullInt32{Int32: 1, Valid: true},
	})
	if err != nil {

		log.Println(err.Error())
	}

	insertedUserID, err := result.LastInsertId()
	if err != nil {
		log.Println(err.Error())

	}
	log.Println(insertedUserID)

	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("hello, world again")
	})
	app.Listen(":3000")

}
