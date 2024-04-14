package database

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
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

func SaveLog(db *sql.DB, logEntry WebsiteLog) error {
	query := `INSERT INTO public.website_monitor_log (url, error,  thread_id, session_id, "timestamp")
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`
	_, err := db.Exec(query, logEntry.URL, logEntry.Error, logEntry.ThreadId, logEntry.SessionId, time.Now())
	if err != nil {
		return fmt.Errorf("error inserting log entry: %v", err)
	}
	return nil
}
