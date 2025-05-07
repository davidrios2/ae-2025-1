package _2_LSP

import "fmt"

// Notifier interface defines the contract for sending notifications.
type Notifier interface {
	Send(message string) error
}

type SoundNotifierInterface interface {
	Notifier // Embed the basic Notifier interface
	PlaySound() error
}

// EmailNotifier sends notifications via email.
type EmailNotifier struct {
	Recipient string
}

// Send sends an email notification.
func (e EmailNotifier) Send(message string) error {
	fmt.Printf("Sending email to %s: %s\n", e.Recipient, message)
	return nil
}

// SMSNotifier sends notifications via SMS.
type SMSNotifier struct {
	PhoneNumber string
}

// Send sends an SMS notification.
func (s SMSNotifier) Send(message string) error {
	fmt.Printf("Sending SMS to %s: %s\n", s.PhoneNumber, message)
	return nil
}

// PushNotifier sends notifications via push notifications.
type PushNotifier struct {
	DeviceToken string
	HasSound    bool
	SoundFile   string
}

// Send sends a push notification.
func (p PushNotifier) Send(message string) error {
	fmt.Printf("Sending push notification to %s: %s\n", p.DeviceToken, message)
	//  DO NOT call p.PlaySound() here.  That violates 02_LSP.
	return nil
}

// PlaySound plays the sound for the push notification.
func (p PushNotifier) PlaySound() error {
	fmt.Printf("Playing sound: %s\n", p.SoundFile)
	return nil
}

// NotificationService manages sending notifications using different notifiers.
type NotificationService struct {
	notifier Notifier
}

// NewNotificationService creates a new NotificationService.
func NewNotificationService(notifier Notifier) *NotificationService {
	return &NotificationService{notifier: notifier}
}

// SendNotification sends a notification using the configured notifier.
func (ns *NotificationService) SendNotification(message string) error {
	return ns.notifier.Send(message)
}

func LspC() {
	emailNotifier := EmailNotifier{Recipient: "user@example.com"}
	smsNotifier := SMSNotifier{PhoneNumber: "+1234567890"}
	pushNotifier := PushNotifier{DeviceToken: "device-token-123", HasSound: true, SoundFile: "sounds.mp3"}

	emailService := NewNotificationService(emailNotifier)
	smsService := NewNotificationService(smsNotifier)
	pushService := NewNotificationService(pushNotifier)

	err := emailService.SendNotification("Welcome to our platform!")
	if err != nil {
		fmt.Println("Error sending email:", err)
	}
	err = smsService.SendNotification("Your order has been shipped.")
	if err != nil {
		fmt.Println("Error sending SMS:", err)
	}
	err = pushService.SendNotification("New message received!")
	pushService.notifier.(SoundNotifierInterface).PlaySound()
	if err != nil {
		fmt.Println("Error sending push notification:", err)
	}

	// You can easily switch between notifiers without changing the core logic
	// because they all adhere to the Notifier interface.
	emailService.notifier = pushNotifier
	emailService.SendNotification(" switched to sms ")

	/*if soundNotifier, ok := pushService.notifier.(SoundNotifierInterface); ok {
		err := soundNotifier.PlaySound()
		if err != nil {
			fmt.Println("Error playing sound:", err)
		}
	}*/

}
