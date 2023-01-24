package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/jackc/pgx/v5"
	_ "github.com/jackc/pgx/v5/stdlib"
)

func main() {
	// Connect to the database
	conn, err := sql.Open("pgx", "host=localhost port=5432 dbname=blog_db user=postgres password=postgres")
	if err != nil {
		log.Fatalf(fmt.Sprintf("Couldn't connect to the database : %v\n", err))
	}
	defer conn.Close()

	// ping the database. check if it response.
	err = conn.Ping()
	if err != nil {
		log.Fatalf("Couldn't ping the database : %v\n", err)
	}
}
