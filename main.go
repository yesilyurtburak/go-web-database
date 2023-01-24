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

// This function inserts new record to the table.
func insertNewUser(conn *sql.DB, name string, email string, password string, uType int) error {
	// create a query string
	query := fmt.Sprintf(`INSERT INTO users (name, email, password, account_created, last_login, user_type) VALUES ('%s', '%s', '%s', current_timestamp, current_timestamp, %d)`, name, email, password, uType)

	_, err := conn.Exec(query) // executes a query.
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

// This function retrieves a record from the table by using its id number.
func getUserData(conn *sql.DB, id int) error {
	var name, email, password, uType string // variables that will store the retrieved data
	// create a query string
	query := fmt.Sprintf(`SELECT id, name, email, password, user_type FROM users WHERE id = %d`, id)
	// QueryRow() is expected to return at most 1 row.
	row := conn.QueryRow(query)
	// Scan() copies the retrieved data into destination.
	err := row.Scan(&id, &name, &email, &password, &uType)

	if err != nil {
		log.Fatal(err)
		return err
	}
	fmt.Println("ID:", id)
	fmt.Println("Name:", name)
	fmt.Println("Email:", email)
	fmt.Println("Password:", password)
	fmt.Println("User Type:", uType)
	return nil
}

// This function updates the user email on the table.
func updateUserEmail(conn *sql.DB, newEmail string, id int) error {
	query := fmt.Sprintf(`UPDATE users SET email = '%s' WHERE id = %d`, newEmail, id)

	_, err := conn.Exec(query)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

// This function deletes a record by its id.
func deleteUserById(conn *sql.DB, id int) error {
	query := fmt.Sprintf(`DELETE FROM users WHERE id = %d`, id)
	_, err := conn.Exec(query)
	if err != nil {
		log.Fatal(err)
		return err
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

	// function call to insert user to database
	err = insertNewUser(conn, "Sally Smith", "ss@gmail.com", "weakpassword", 3)
	if err != nil {
		log.Fatal(err)
	}

	err = getAllRowData(conn) // check the table after inserting data.
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("-------------------------")

	// Get a user data from the database by its id
	getUserData(conn, 1)

	fmt.Println("-------------------------")

	// Update the email of a user.
	err = updateUserEmail(conn, "yesilyurtburak@google.com", 1)
	if err != nil {
		log.Fatal(err)
	}
	getUserData(conn, 1) // check the spesific user's data.

	fmt.Println("-------------------------")

	// delete the user by its id
	err = deleteUserById(conn, 1)
	if err != nil {
		log.Fatal(err)
	}
	err = getAllRowData(conn) // check all information in the table
	if err != nil {
		log.Fatal(err)
	}
}
