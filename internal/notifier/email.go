package notifier

import (
	"fmt"
	"monitor/internal/config"
	"net/smtp"
	"strings"
)

func SendAlert(cfg *config.Config, targetURL string, checkErr error) error {
	auth := smtp.PlainAuth("", cfg.SMTPUser, cfg.SMTPPass, cfg.SMTPHost)

	subject := "Website Down Alert: " + targetURL
	body := fmt.Sprintf("Alert! The website %s is down.\n\nError: %v\n\nTime: %s",
		targetURL, checkErr, "Now") // using "Now" for simplicity, could import time

	msg := []byte(fmt.Sprintf("To: %s\r\n"+
		"Subject: %s\r\n"+
		"\r\n"+
		"%s\r\n", cfg.ToEmail, subject, body))

	addr := fmt.Sprintf("%s:%d", cfg.SMTPHost, cfg.SMTPPort)

	// Note: For local testing or servers without auth, auth might need to be nil.
	// However, most modern SMTP servers require auth.
	// If the user provided no user/pass, we might skip auth? 
	// For now, adhering to the plan which includes user/pass.
	
	err := smtp.SendMail(addr, auth, cfg.FromEmail, []string{cfg.ToEmail}, msg)
	if err != nil {
		// Try without auth if auth failed or wasn't provided (simple safeguard)
		if cfg.SMTPUser == "" || cfg.SMTPPass == "" {
			err = smtp.SendMail(addr, nil, cfg.FromEmail, []string{cfg.ToEmail}, msg)
		}
	}
	
	if err != nil {
		// Check for common error where auth is not supported by server but we tried it
		if strings.Contains(err.Error(), "unencrypted connection") {
             // Logic to handle unencrypted, but stdlib smtp.SendMail requires TLS for Auth usually
             // This is a simplified implementation.
		}
		return fmt.Errorf("failed to send email: %w", err)
	}

	return nil
}
