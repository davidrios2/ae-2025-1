package _4_DIP

import "fmt"

// EmailService is a low-level module for sending emails.
type EmailService struct {
	smtpServer string
}

func (e *EmailService) SendEmail(recipient, message string) {
	fmt.Printf("Sending email to %s via %s: %s\n", recipient, e.smtpServer, message)
}

// NotificationService is a high-level module that directly depends on EmailService.
type NotificationService struct {
	emailService *EmailService
}

func (n *NotificationService) SendNotification(recipient, message string) {
	n.emailService.SendEmail(recipient, message)
}

func DipI() {
	emailService := &EmailService{smtpServer: "smtp.example.com"}
	notificationService := &NotificationService{emailService: emailService}

	notificationService.SendNotification("user@example.com", "Hello, user!")
}
