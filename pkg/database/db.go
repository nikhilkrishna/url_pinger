package database

import (
	"database/sql"
	_ "fmt"
	_ "github.com/lib/pq"
	"log"
)

func InitDB(dbConn string) *sql.DB {
	db, err := sql.Open("postgres", dbConn)
	if err != nil {
		log.Fatalf("Error opening database connection: %v", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("Cannot ping the database: %v", err)
	}

	return db
}

type WebsiteLog struct {
	SessionId string
	ThreadId  int
	URL       string
	Response  string
	Error     string
	Pattern   string
}
