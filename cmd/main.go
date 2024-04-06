package main

import (
	"log"

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
}
