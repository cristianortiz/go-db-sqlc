package middlewares

import (
	"strconv"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

//IsAuthenticated middleware functions checks if jwt existis stored in fiber cookie and validates user session
func IsAuthenticated(c *fiber.Ctx) error {
	//get the cookie from fiber context
	cookie := c.Cookies("jwt")
	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("ReactGoSqlc"), nil
	})
	if err != nil || !token.Valid {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"msg": "Invalid credentials",
		})
	}

	return c.Next()
}

//GetUserIdFromJWT func returns the id from a logged user from jwt claims stored in fiber cookie
func GetUserIdFromJWT(c *fiber.Ctx) (int64, error) {
	cookie := c.Cookies("jwt")
	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("ReactGoSqlc"), nil

	})
	if err != nil {
		return 0, err
	}

	payload := token.Claims.(*jwt.StandardClaims)
	id, _ := strconv.ParseInt(payload.Subject, 0, 64)
	return id, nil

}
