package controllers

import (
	"context"
	"go-db-sqlc/src/database"
	"go-db-sqlc/src/database/products"
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func GetProducts(c *fiber.Ctx) error {
	//sqlc context and queries struct
	ctx := context.Background()
	queries := products.New(database.DB)
	result, err := queries.GetAllProducts(ctx)
	if err != nil {
		return err
	}

	return c.JSON(result)

}

func CreateProduct(c *fiber.Ctx) error {
	//sqlc context and queries struct
	ctx := context.Background()
	queries := products.New(database.DB)

	//get data for the request
	var data map[string]string
	//get data from the http request and assign it to data map
	err := c.BodyParser(&data)
	if err != nil {
		return err
	}
	p, _ := strconv.ParseFloat(data["price"], 64)
	// create an user in DB
	result, err := queries.CreateNewProduct(ctx, products.CreateNewProductParams{
		Title:       data["title"],
		Description: data["description"],
		Image:       data["image"],
		Price:       p,
	})
	if err != nil {

		log.Println(err.Error())
	}

	insertedProductID, err := result.LastInsertId()
	if err != nil {
		log.Println(err.Error())

	}
	log.Println(insertedProductID)

	return c.JSON(fiber.Map{
		"newProductID": insertedProductID,
	})

}
