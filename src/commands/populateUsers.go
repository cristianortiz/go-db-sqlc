package main

import (
	"context"
	"database/sql"
	"go-db-sqlc/src/database"
	"go-db-sqlc/src/database/users"
	"go-db-sqlc/src/utils"
	"log"

	"github.com/bxcodec/faker/v3"
)

func main() {
	database.DBConnect()
	//sqlc context and queries struct
	ctx := context.Background()
	queries := users.New(database.DB)
	var pass string

	for i := 0; i < 30; i++ {
		pass, _ = utils.SetAndEncryptPassword("1234")
		// create an user in DB
		result, _ := queries.CreateUser(ctx, users.CreateUserParams{
			Firstname:    faker.FirstName(),
			Lastname:     faker.LastName(),
			Email:        faker.Email(),
			Upassword:    pass,
			Isambassador: sql.NullInt32{Int32: 1, Valid: true},
		})
		log.Println(result.LastInsertId())

	}
}
