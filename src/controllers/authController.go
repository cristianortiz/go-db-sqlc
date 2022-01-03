package controllers

import (
	"context"
	"database/sql"
	"go-db-sqlc/src/database"
	"go-db-sqlc/src/database/users"
	"go-db-sqlc/src/middlewares"
	"go-db-sqlc/src/utils"
	"time"

	"log"

	"github.com/gofiber/fiber/v2"
)

//controller function register endpoint to register a new user account
func Register(c *fiber.Ctx) error {
	//sqlc context and queries struct
	ctx := context.Background()
	queries := users.New(database.DB)

	//get data for the request
	var data map[string]string
	//get data from the http request and assign it to data map
	err := c.BodyParser(&data)
	if err != nil {
		return err
	}
	//check if the email is already in use
	used, _ := queries.UserEmailExists(ctx, data["email"])
	if used == 1 {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"msg": "Username is alreay in use",
		})
	}
	//validations
	if data["upassword"] != data["upassword_confirm"] {
		return c.JSON(fiber.Map{
			"message": "passwords do not match",
		})
	}
	password, err := utils.SetAndEncryptPassword(data["upassword"])
	if err != nil {
		return err
	}

	// create an user in DB
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

//Login function checks if an user exists and verify their password
func Login(c *fiber.Ctx) error {
	//sqlc context and queries struct
	ctx := context.Background()
	queries := users.New(database.DB)

	//map to store the request data
	var data map[string]string
	//get data from the http request and assign it to data map
	err := c.BodyParser(&data)
	if err != nil {
		return err
	}
	//check if the email is already in use
	founded, _ := queries.UserEmailExists(ctx, data["email"])

	if founded == 0 {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"msg": "Username or password are wrong",
		})
	}
	user, _ := queries.GetUserByEmail(ctx, data["email"])
	//bcrypt only works with slice of bytes data,hash the password received as parameter
	//and the pass returned by the DB
	err = utils.ComparePassword(data["upassword"], user)
	//if pass are not equals  response the error
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"msg": "Username or password are wrong",
		})
	}
	//generates JWT for use auth
	token, err := utils.GeneratesJWT(user)
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"msg": "Invalid Credentials",
		})
	}
	//the JWT recorded in a cookie
	//cookie expiration time
	expirationTime := time.Now().Add(24 * time.Hour)
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  expirationTime,
		HTTPOnly: true, //to be sent to the backend
	}
	//set cookie in fiber context
	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"msg:": "success",
	})
}

//AuthenticatedUser returns the data of a logged user, using the jwt in cookie stored in fiber context
func GetUser(c *fiber.Ctx) error {
	//sqlc context and queries struct
	ctx := context.Background()
	queries := users.New(database.DB)

	id, _ := middlewares.GetUserIdFromJWT(c)
	//payload.Subject stores the id of the user logged
	user, _ := queries.GetUserParamsByID(ctx, id)
	return c.JSON(user)

}

//UpdateUserInfo() controller function to update the data of an existing an logged user
func UpdateUserInfo(c *fiber.Ctx) error {
	//sqlc context and queries struct
	ctx := context.Background()
	queries := users.New(database.DB)

	//map to store the thata received in request
	var data map[string]string
	//get data from the http request and assign it to data map
	err := c.BodyParser(&data)
	if err != nil {
		return err
	}
	//get logged suer id from jwt claims
	id, _ := middlewares.GetUserIdFromJWT(c)

	//updates a logged user basic data in DB
	_, err = queries.UpdateUserInfo(ctx, users.UpdateUserInfoParams{
		Firstname: data["firstname"],
		Lastname:  data["lastname"],
		Email:     data["email"],
		ID:        id,
	})
	if err != nil {

		log.Println(err.Error())
	}

	user, _ := queries.GetUserParamsByID(ctx, id)
	return c.JSON(user)

}

//UpdateUserInfo() controller function to update the data of an existing an logged user
func UpdateUserPassword(c *fiber.Ctx) error {
	//sqlc context and queries struct
	ctx := context.Background()
	queries := users.New(database.DB)
	//map to store the thata received in request
	var data map[string]string
	//get data from the http request and assign it to data map
	err := c.BodyParser(&data)
	if err != nil {
		return err
	}

	if data["upassword"] != data["upassword_confirm"] {
		return c.JSON(fiber.Map{
			"message": "passwords do not match",
		})
	}
	password, err := utils.SetAndEncryptPassword(data["upassword"])
	if err != nil {
		return err
	}
	id, _ := middlewares.GetUserIdFromJWT(c)
	//updates a logged user basic data in DB
	_, err = queries.UpdateUserPassword(ctx, users.UpdateUserPasswordParams{
		Upassword: password,
		ID:        id,
	})
	if err != nil {

		log.Println(err.Error())
	}

	return c.JSON(fiber.Map{
		"msg": "password updated",
	})

}

//Logout function reset the the jwt in cookie to invalidate user credentials
func Logout(c *fiber.Ctx) error {
	cookie := fiber.Cookie{
		Name:  "jwt",
		Value: "",
		//make the cookie already expired one hour ago
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
	}
	c.Cookie(&cookie)
	return c.JSON(fiber.Map{
		"msg": "user logout",
	})
}
