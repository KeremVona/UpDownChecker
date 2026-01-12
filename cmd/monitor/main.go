package main

import (
	"flag"
	"fmt"
	"log"
	"monitor/internal/config"
	"monitor/internal/monitor"
	"os"
)

func waitExit() {
	fmt.Println("\nPress Enter to exit...")
	fmt.Scanln()
	os.Exit(1)
}

func main() {
	sitesFile := flag.String("sites", "sites.csv", "Path to the sites CSV file")
	configFile := flag.String("config", "config.json", "Path to the configuration JSON file")
	flag.Parse()

	// Load Config
	cfg, err := config.LoadConfig(*configFile)
	if err != nil {
		log.Printf("Error loading config: %v", err)
		log.Println("Make sure config.json exists and is valid.")
		waitExit()
	}

	// Load Targets
	targets, err := config.LoadTargets(*sitesFile)
	if err != nil {
		log.Printf("Error loading targets: %v", err)
		log.Println("Make sure sites.csv exists and is valid.")
		waitExit()
	}

	if len(targets) == 0 {
		log.Println("No targets found in CSV.")
		waitExit()
	}

	log.Printf("Loaded %d targets. Starting monitor...", len(targets))

	// Start Monitoring
	monitor.Start(cfg, targets)
}
