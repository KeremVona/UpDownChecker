package monitor

import (
	"log"
	"monitor/internal/checker"
	"monitor/internal/config"
	"monitor/internal/notifier"
	"sync"
	"time"
)

// Start begins monitoring the targets.
func Start(cfg *config.Config, targets []config.Target) {
	var wg sync.WaitGroup

	for _, t := range targets {
		wg.Add(1)
		go func(target config.Target) {
			defer wg.Done()
			monitorTarget(cfg, target)
		}(t)
	}

	wg.Wait()
}

func monitorTarget(cfg *config.Config, target config.Target) {
	log.Printf("Starting monitor for %s with interval %v", target.URL, target.Interval)
	ticker := time.NewTicker(target.Interval)
	defer ticker.Stop()

	// Initial check
	checkAndNotify(cfg, target)

	for range ticker.C {
		checkAndNotify(cfg, target)
	}
}

func checkAndNotify(cfg *config.Config, target config.Target) {
	err := checker.Check(target.URL)
	if err != nil {
		log.Printf("[DOWN] %s - Error: %v", target.URL, err)
		if sendErr := notifier.SendAlert(cfg, target.URL, target.ToEmail, err); sendErr != nil {
			log.Printf("Failed to send alert for %s: %v", target.URL, sendErr)
		} else {
			log.Printf("Alert sent for %s", target.URL)
		}

		if target.ToTelegram != "" {
			if sendErr := notifier.SendTelegramAlert(cfg, target.URL, target.ToTelegram, err); sendErr != nil {
				log.Printf("Failed to send telegram alert for %s: %v", target.URL, sendErr)
			} else {
				log.Printf("Telegram alert sent for %s", target.URL)
			}
		}
	} else {
		log.Printf("[UP] %s", target.URL)
	}
}
