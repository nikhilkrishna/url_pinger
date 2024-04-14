package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"url_pinger/pkg/config"
	"url_pinger/pkg/database"
	"url_pinger/pkg/logger"
	"url_pinger/pkg/monitor"
)

func main() {

	// Initialize Logger
	log := logger.NewStdoutLogger()
	defer log.Close()

	cfg, err := config.LoadConfig(".env")
	if err != nil {
		log.Log("Error loading the .env file")
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

	wm := monitor.NewWebsiteMonitor(configs, *sessionID, log, db)
	wm.Start()

	// Wait for shutdown signal
	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, os.Interrupt, syscall.SIGTERM)
	<-stopChan
	log.Log("Shutting down website monitoring...")

}
