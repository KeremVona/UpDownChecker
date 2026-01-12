package checker

import (
	"fmt"
	"net/http"
	"time"
)

// Check performs a HEAD request to the target URL to check availability.
func Check(url string) error {
	client := http.Client{
		Timeout: 10 * time.Second,
	}

	resp, err := client.Head(url)
	if err != nil {
		return fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		return fmt.Errorf("status code %d", resp.StatusCode)
	}

	return nil
}
