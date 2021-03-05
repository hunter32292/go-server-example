package dao

import (
	"database/sql"
	"log"

	// Imported for the connection keep alive on the database function
	_ "github.com/lib/pq"
)

// Database - The Connection object to a backend storage database
var Database *sql.DB

// CreateConnection - Start Connection to database and store the connection in the dao package
func CreateConnection() {
	db, err := sql.Open("postgres", "user=postgres host=localhost dbname=postgres sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	Database = db
}

// Close - close the database connection
func Close() {
	err := Database.Close()
	if err != nil {
		log.Fatal(err)
	}
}

// CheckConnection - Test the database connection health
func CheckConnection() (bool, error) {
	err := Database.Ping()
	if err != nil {
		return false, err
	} else {
		return true, nil
	}
}
