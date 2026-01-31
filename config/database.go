package config

import (
	"log"

	_ "github.com/lib/pq"
	"database/sql"
)

// InitDB initializes and returns a database connection
func InitDB(connectionString string) (*sql.DB, error) {
	// Open a database connection
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}

	// Verify the connection
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	// set connection pool settings if needed
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)

	log.Println("Database connection established")
	return db, nil
}