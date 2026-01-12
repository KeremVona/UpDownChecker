package main

import (
	"flag"
	"log"
	"monitor/internal/config"
	"monitor/internal/monitor"
	"os"
)

func main() {
	sitesFile := flag.String("sites", "sites.csv", "Path to the sites CSV file")
	configFile := flag.String("config", "config.json", "Path to the configuration JSON file")
	flag.Parse()

	// Load Config
	cfg, err := config.LoadConfig(*configFile)
	if err != nil {
		log.Printf("Error loading config: %v", err)
		log.Println("Make sure config.json exists and is valid.")
		os.Exit(1)
	}

	// Load Targets
	targets, err := config.LoadTargets(*sitesFile)
	if err != nil {
		log.Printf("Error loading targets: %v", err)
		log.Println("Make sure sites.csv exists and is valid.")
		os.Exit(1)
	}

	if len(targets) == 0 {
		log.Println("No targets found in CSV.")
		os.Exit(1)
	}

	log.Printf("Loaded %d targets. Starting monitor...", len(targets))

	// Start Monitoring
	monitor.Start(cfg, targets)
}
