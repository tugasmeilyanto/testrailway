package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// Replace with your Railway MySQL database credentials
	dsn := "mysql://root:QQnjRRqLtuoHIbDUMUojCpHibYCWfMTg@monorail.proxy.rlwy.net:55482/railway"

	// Open a connection to the database
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Error opening database: %v\n", err)
	}
	defer db.Close()

	// Verify the connection
	err = db.Ping()
	if err != nil {
		log.Fatalf("Error pinging database: %v\n", err)
	}

	fmt.Println("Successfully connected to the database!")

	// Perform a simple query
	var version string
	err = db.QueryRow("SELECT VERSION()").Scan(&version)
	if err != nil {
		log.Fatalf("Error executing query: %v\n", err)
	}

	fmt.Printf("Database version: %s\n", version)
}
