package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"url_pinger/pkg/config"
	"url_pinger/pkg/database"
)

func main() {
	cfg, err := config.LoadConfig(".env")
	if err != nil {
		log.Fatal("Error loading the .env file")
	}

	db := database.InitDB(cfg.DBConn)
	defer db.Close()

	csvFilePath := flag.String("csv", "websites.csv", "Path to the CSV file with website configs")
	sessionID := flag.String("session", "default-session", "Session ID for logging")
	flag.Parse()

	configs, err := config.LoadWebsiteSettings(*csvFilePath)
	if err != nil {
		fmt.Println("Error loading website configurations:", err)
		os.Exit(1)
	}

}
