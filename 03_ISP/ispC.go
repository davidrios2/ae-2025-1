package _3_ISP

import "fmt"

// PaymentProcessorInterface interface for basic payment processing.  <- Changed name
type PaymentProcessorInterface interface {
	ProcessPayment(amount float64) error
	RefundPayment(transactionID string, amount float64) error
}

// AuthCaptureInterface interface for authorization and capture.  <- Changed name
type AuthCaptureInterface interface {
	CapturePayment(transactionID string) error
	VoidPayment(transactionID string) error
}

// StripeProcessor implements both PaymentProcessorInterface and AuthCaptureInterface.
type StripeProcessorC struct{}

func (s StripeProcessorC) ProcessPayment(amount float64) error {
	fmt.Println("PaymentProcessorInterface.ProcessPayment:", amount)
	return nil
}

func (s StripeProcessorC) RefundPayment(transactionID string, amount float64) error {
	fmt.Println("PaymentProcessorInterface.RefundPayment:", transactionID, amount)
	return nil
}

func (s StripeProcessorC) CapturePayment(transactionID string) error {
	fmt.Println("AuthCaptureInterface.CapturePayment:", transactionID)
	return nil
}

func (s StripeProcessorC) VoidPayment(transactionID string) error {
	fmt.Println("AuthCaptureInterface.VoidPayment:", transactionID)
	return nil
}

// PayPalProcessor implements PaymentProcessorInterface.
type PayPalProcessorC struct{}

func (p PayPalProcessorC) ProcessPayment(amount float64) error {
	fmt.Println("PaymentProcessorInterface.ProcessPayment:", amount)
	return nil
}

func (p PayPalProcessorC) RefundPayment(transactionID string, amount float64) error {
	fmt.Println("PaymentProcessorInterface.RefundPayment:", transactionID, amount)
	return nil
}

func IspC() {
	stripe := StripeProcessorC{}
	payPal := PayPalProcessorC{}

	stripe.ProcessPayment(100.00)
	stripe.RefundPayment("123", 50.00)
	stripe.CapturePayment("123")
	stripe.VoidPayment("123")

	payPal.ProcessPayment(50.00)
	payPal.RefundPayment("456", 20.00)
}
