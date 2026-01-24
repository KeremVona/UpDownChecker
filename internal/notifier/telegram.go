package notifier

import (
	"bytes"
	"encoding/json"
	"fmt"
	"monitor/internal/config"
	"net/http"
	"time"
)

type telegramMessage struct {
	ChatID string `json:"chat_id"`
	Text   string `json:"text"`
}

func SendTelegramAlert(cfg *config.Config, targetURL string, chatID string, checkErr error) error {
	if cfg.TelegramToken == "" {
		return fmt.Errorf("telegram token not provided")
	}

	message := fmt.Sprintf("Alert! The website %s is down.\n\nError: %v\n\nTime: %s",
		targetURL, checkErr, time.Now().Format(time.RFC1123))

	payload := telegramMessage{
		ChatID: chatID,
		Text:   message,
	}

	body, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	url := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", cfg.TelegramToken)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to send telegram message, status: %s", resp.Status)
	}

	return nil
}
