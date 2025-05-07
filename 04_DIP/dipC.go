package _4_DIP

import "fmt"

// MessageSender is an interface for sending messages.  (Abstraction)
type MessageSender interface {
	SendMessage(recipient, message string)
}

// MailService is a low-level module for sending emails.
type MailService struct {
	smtpServer string
}

func (e *MailService) SendMessage(recipient, message string) {
	fmt.Printf("Sending email to %s via %s: %s\n", recipient, e.smtpServer, message)
	// Code to send email using e.smtpServer
}

// SMS service is a low-level module for sending SMS.
type SMSService struct {
	twilioAccountSid string
	twilioAuthToken  string
}

func (s *SMSService) SendMessage(recipient, message string) {
	fmt.Printf("Sending SMS to %s via Twilio: %s\n", recipient, message)
}

// NtfService is a high-level module that depends on the MessageSender abstraction.
type NtfService struct {
	messageSender MessageSender
}

func (n *NtfService) SendNotification(recipient, message string) {
	n.messageSender.SendMessage(recipient, message)
}

func DipC() {
	emailService := &MailService{smtpServer: "smtp.example.com"}
	smsService := &SMSService{twilioAccountSid: "ACxxxxxxxx", twilioAuthToken: "your_auth_token"}

	notificationServiceWithEmail := &NtfService{messageSender: emailService}
	notificationServiceWithSMS := &NtfService{messageSender: smsService}

	notificationServiceWithEmail.SendNotification("user@example.com", "Hello, user!")
	notificationServiceWithSMS.SendNotification("+1234567890", "Your order is ready.")
}
