package _2_LSP

import "fmt"

// NotifierInterface defines the contract for sending notifications.
type NotifierInterface interface {
	Transmit(message string) error
}

// EmailMessenger sends notifications via email.
type EmailMessenger struct {
	Receiver string
}

// Transmit sends an email notification.
func (e EmailMessenger) Transmit(message string) error {
	fmt.Printf("Sending email to %s: %s\n", e.Receiver, message)
	return nil
}

// SmsMessenger sends notifications via SMS.
type SmsMessenger struct {
	MobileNumber string
}

// Transmit sends an SMS notification.
func (s SmsMessenger) Transmit(message string) error {
	fmt.Printf("Sending SMS to %s: %s\n", s.MobileNumber, message)
	return nil
}

// PushMessenger sends notifications via push notifications.
type PushMessenger struct {
	DeviceID string
	hasSound bool
	Sound    string
}

// Transmit sends a push notification.
func (p PushMessenger) Transmit(message string) error {
	fmt.Printf("Sending push notification to %s: %s\n", p.DeviceID, message)

	if p.hasSound { // This check makes PushMessenger behave differently.
		fmt.Println("...with sound!")
	}
	return nil
}

// NotificationManager manages sending notifications using different notifiers.
type NotificationManager struct {
	notifier NotifierInterface
}

// NewNotificationManager creates a new NotificationManager.
func NewNotificationManager(notifier NotifierInterface) *NotificationManager {
	return &NotificationManager{notifier: notifier}
}

// SendNotification sends a notification using the configured notifier.
func (ns *NotificationManager) SendNotification(message string) error {
	return ns.notifier.Transmit(message)
}

func LspI() {
	emailNotifier := EmailMessenger{Receiver: "user@example.com"}
	smsNotifier := SmsMessenger{MobileNumber: "+1234567890"}
	pushNotifier := PushMessenger{DeviceID: "device-token-123", hasSound: true, Sound: "sound.mp3"}

	emailService := NewNotificationManager(emailNotifier)
	smsService := NewNotificationManager(smsNotifier)
	pushService := NewNotificationManager(pushNotifier)

	err := emailService.SendNotification("Welcome to our platform!")
	if err != nil {
		fmt.Println("Error sending email:", err)
	}
	err = smsService.SendNotification("Your order has been shipped.")
	if err != nil {
		fmt.Println("Error sending SMS:", err)
	}
	err = pushService.SendNotification("New message received!")
	if err != nil {
		fmt.Println("Error sending push notification:", err)
	}

	// You can easily switch between notifiers without changing the core logic
	// because they all adhere to the NotifierInterface interface.
	emailService.notifier = pushNotifier
	emailService.SendNotification("switched to push ")
}
