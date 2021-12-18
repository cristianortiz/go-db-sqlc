package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func DBConnect() {
	var err error
	DB, err = sql.Open("mysql", "root:invernalia2013@tcp(127.0.0.1:3306)/ambassador")
	//defer DB.Close()
	if err != nil {
		log.Fatal(err)
	}

	var version string

	err2 := DB.QueryRow("SELECT VERSION()").Scan(&version)

	if err2 != nil {
		log.Fatal(err2)
	}

	fmt.Println("DB is Running on air : " + version)
	//return DB

}
