package database

import (
	"database/sql"
	"log"
)

func Connect(connectionString string) *sql.DB {
	db, err := sql.Open("postgres", connectionString)
    if err != nil {
        log.Fatalf("Error connecting to database: %v", err)
    }

	if err := db.Ping(); err != nil {
		log.Fatalf("Database is not reachable: %v", err)
	}

	log.Println("Successfully connected to the database")

    return db
}