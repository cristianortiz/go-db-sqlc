package controllers

import (
	"context"
	"go-db-sqlc/src/database"
	"go-db-sqlc/src/database/users"

	"github.com/gofiber/fiber/v2"
)

func GetAmbassadors(c *fiber.Ctx) error {
	//sqlc context and queries struct
	ctx := context.Background()
	queries := users.New(database.DB)
	result, err := queries.GetAmbassadors(ctx)
	if err != nil {
		return err
	}

	return c.JSON(result)
}
