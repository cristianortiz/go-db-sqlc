package controllers

import (
	"context"
	"database/sql"
	"go-db-sqlc/src/database"
	"go-db-sqlc/src/database/users"

	"log"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *fiber.Ctx) error {
	//get data for the request
	var data map[string]string
	//get data from the http request and assign it to data map
	err := c.BodyParser(&data)
	if err != nil {
		return err
	}
	//validations
	if data["upassword"] != data["upassword_confirm"] {
		return c.JSON(fiber.Map{
			"message": "passwords do not match",
		})
	}
	password, err := PasswordEncryption(data["upassword"])
	if err != nil {
		return err
	}
	//sqlc context and queries struct
	ctx := context.Background()
	queries := users.New(database.DB)

	// create an user
	result, err := queries.CreateUser(ctx, users.CreateUserParams{
		Firstname:    data["firstname"],
		Lastname:     data["lastname"],
		Email:        data["email"],
		Upassword:    password,
		Isambassador: sql.NullInt32{Int32: 0, Valid: true},
	})
	if err != nil {

		log.Println(err.Error())
	}

	insertedUserID, err := result.LastInsertId()
	if err != nil {
		log.Println(err.Error())

	}
	log.Println(insertedUserID)

	return c.JSON(fiber.Map{
		"newUserID": insertedUserID,
	})

}

//PasswordEncryption encrypts the user pass using bcrypt library
func PasswordEncryption(pass string) (string, error) {
	//number of layer for encryption algo
	cost := 8
	//GeneratesFormPassword only accepts a slice of bytes []byte
	bytes, err := bcrypt.GenerateFromPassword([]byte(pass), cost)
	return string(bytes), err
}
