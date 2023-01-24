package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/jackc/pgx/v5"
	_ "github.com/jackc/pgx/v5/stdlib"
)

// This function takes connection as parameter and returns a possible error.
func getAllRowData(conn *sql.DB) error {
	rows, err := conn.Query("SELECT id, name, email FROM users") // sql query
	if err != nil {
		log.Println(err)
		return err
	}
	defer rows.Close() // close after it has done.

	// create variables to hold the scanned data.
	var id int
	var name, email string

	for rows.Next() {
		err := rows.Scan(&id, &name, &email) // scan the rows in the table and copy the values into vars.
		if err != nil {
			log.Println(err)
		}
		fmt.Println("Data : ", id, name, email)
	}
	if err != nil {
		log.Fatal("Error reading data : ", err)
	}
	return nil
}

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

	// function call for fetching data from database
	err = getAllRowData(conn)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("-------------------------")
}
