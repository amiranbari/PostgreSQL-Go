package main

import (
	"database/sql"
	"fmt"
	_ "github.com/jackc/pgx/v4/stdlib"
	"log"
)

var first_name, last_name, email string
var id int

func main() {
	db, err := sql.Open("pgx", "host=localhost port=5432 dbname=test user=postgres password=123456")
	if err != nil {
		log.Fatal("Unable to connect: %v\n", err)
	}
	defer db.Close()

	log.Println("Connected to database!")

	//test my connection
	err = db.Ping()
	if err != nil {
		log.Fatal("Cannot ping to database")
	}

	log.Println("Ping to database successfully!")

	//show all rows database
	err = getAllRows(db)
	if err != nil {
		log.Fatal(err)
	}

	//insert into table users
	query := `insert into users (first_name, last_name, email, created_at, updated_at) values ($1, $2, $3, $4, $5)`
	_, err = db.Exec(query, "Amir", "Anbari", "amiranbari32@gmail.com", "2002-06-15 01:56:30", "2002-06-15 01:56:31")

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Insert into table!")

	//show all rows database
	err = getAllRows(db)
	if err != nil {
		log.Fatal(err)
	}

	//update table users
	query = `update users set first_name = $1 where first_name = $2`
	_, err = db.Exec(query, "Amir2", "Amir")

	if err != nil {
		log.Fatal(err)
	}

	log.Println("update table!")

	//show all rows database
	err = getAllRows(db)
	if err != nil {
		log.Fatal(err)
	}

	query = `select id, first_name, last_name, email from users where id = $1`
	row := db.QueryRow(query, 1)
	err = row.Scan(&id, &first_name, &last_name, &email)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("QueryRow returns", id, first_name, last_name, email)

	query = `delete from users where id = $1`
	_, err = db.Exec(query, 6)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Delete a row!")

	//show all rows database
	err = getAllRows(db)
	if err != nil {
		log.Fatal(err)
	}

}

// getAllRows get all rows from database
func getAllRows(conn *sql.DB) error {
	rows, err := conn.Query("select id, first_name, last_name, email from users")
	if err != nil {
		log.Println(err)
		return err
	}
	defer rows.Close()

	for rows.Next() {
		fmt.Println("----------------------------")

		err := rows.Scan(&id, &first_name, &last_name, &email)
		if err != nil {
			log.Println(err)
			return err
		}
		fmt.Println("Record is", id, first_name, last_name, email)
	}

	if err = rows.Err(); err != nil {
		log.Fatal("Error scan rows", err)
	}

	fmt.Println("----------------------------")

	return nil
}
