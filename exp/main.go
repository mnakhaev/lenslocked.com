package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "qwe123QWE"
	dbname   = "lenslocked_dev"
)

func main() {
	connectionString := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname,
	)

	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("successfully connected!")

	var id int
	row := db.QueryRow(`
		INSERT INTO users(name,email)
		VALUES($1, $2) RETURNING id`,
		"John Doe", "john.doe@test.com",
	)
	err = row.Scan(&id)
	if err != nil {
		panic(err)
	}
	fmt.Printf("id=%d\n", id)
}
