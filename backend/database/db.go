package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func ConnectDB() {
	var err error
	DB, err = sql.Open("postgres", "host=localhost user=postgres password=postgres dbname=taskdb sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to Database!")
}
