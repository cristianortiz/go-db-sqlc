package controllers

import (
	"context"
	"go-db-sqlc/src/database"
	"go-db-sqlc/src/database/products"
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

//GetProducts controller function retrives the whole list of products from DB
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

//GetProductByID controller function to retrieve the info of a specific product
func GetProductByID(c *fiber.Ctx) error {
	//sqlc context and queries struct
	ctx := context.Background()
	queries := products.New(database.DB)
	//id in http request is a string, cast to the int64 type for query to DB
	id, _ := strconv.ParseInt(c.Params("id"), 0, 64)
	result, err := queries.GetProductByID(ctx, id)
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
		Category:    data["category"],
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

	return c.JSON(fiber.Map{
		"newProductID": insertedProductID,
	})

}

func UpdateProduct(c *fiber.Ctx) error {
	//sqlc context and queries struct
	ctx := context.Background()
	queries := products.New(database.DB)

	//get data map for store the http request data
	var data map[string]string
	//get data from the http request and assign it to data map
	err := c.BodyParser(&data)
	if err != nil {
		return err
	}
	//cast string id from http request data into int64
	id, _ := strconv.ParseInt(c.Params("id"), 0, 64)
	//cast the price in string format to float
	p, _ := strconv.ParseFloat(data["price"], 64)

	// update the product info in DB
	_, err = queries.UpdateProduct(ctx, products.UpdateProductParams{
		ID:          id,
		Category:    data["category"],
		Title:       data["title"],
		Description: data["description"],
		Image:       data["image"],
		Price:       p,
	})
	if err != nil {

		log.Println(err.Error())
	}

	product, _ := queries.GetProductByID(ctx, id)
	return c.JSON(product)

}
func DeleteProduct(c *fiber.Ctx) error {
	//sqlc context and queries struct
	ctx := context.Background()
	queries := products.New(database.DB)

	//cast string id from http request data into int64
	id, _ := strconv.ParseInt(c.Params("id"), 0, 64)

	// update the product info in DB
	result, err := queries.DeleteProduct(ctx, id)
	if err != nil {

		log.Println(err.Error())
	}

	deletedProduct, err := result.RowsAffected()
	if err != nil {
		log.Println(err.Error())

	}
	return c.JSON(deletedProduct)

}
