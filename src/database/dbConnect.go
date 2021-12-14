package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func DBConnect() {
	db, err := sql.Open("mysql", "root:invernalia2013@tcp(127.0.0.1:3306)/ambassador")
	defer db.Close()

	if err != nil {
		log.Fatal(err)
	}

	var version string

	err2 := db.QueryRow("SELECT VERSION()").Scan(&version)

	if err2 != nil {
		log.Fatal(err2)
	}

	fmt.Println("DB is Running : " + version)

}
