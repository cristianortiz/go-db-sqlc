package utils

import (
	"go-db-sqlc/src/database/users"
	"strconv"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

//PasswordEncryption encrypts the user pass using bcrypt library
func SetAndEncryptPassword(pass string) (string, error) {
	//number of layer for encryption algo
	cost := 8
	//GeneratesFormPassword only accepts a slice of bytes []byte
	bytes, err := bcrypt.GenerateFromPassword([]byte(pass), cost)
	return string(bytes), err
}

//comparePassword receiver checks user hashed pass in DB wit the user pass usin gbcrypt library
func ComparePassword(pass string, t users.User) error {
	//bcrypt only works with slice of bytes data,hash the password received as parameter
	//and the pass returned by the DB
	return bcrypt.CompareHashAndPassword([]byte(t.Upassword), []byte(pass))
}

//GeneratesJWT receives a models.user object and create the JWT for user auth
func GeneratesJWT(t users.User) (string, error) {
	//the jwt token is an slice of bytes
	key := []byte("ReactGoSqlc")
	//claims (privileges) section of the token to add in paylod
	payload := jwt.StandardClaims{
		Subject:   strconv.Itoa(int(t.ID)),
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
	}
	//header part of token, encrypton algo
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	//sign the token with key slice of bytes
	tokenStr, err := token.SignedString(key)
	if err != nil {
		return tokenStr, err
	}
	return tokenStr, nil

}
