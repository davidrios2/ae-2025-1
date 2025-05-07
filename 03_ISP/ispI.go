package _3_ISP

import "fmt"

// PaymentProcessor interface defines all payment operations.
type PaymentProcessor interface {
	ProcessPayment(amount float64) error
	RefundPayment(transactionID string, amount float64) error
	CapturePayment(transactionID string) error // For auth & capture
	VoidPayment(transactionID string) error    // Cancels the authorization
}

// StripeProcessor implements PaymentProcessor.
type StripeProcessor struct{}

func (s StripeProcessor) ProcessPayment(amount float64) error {
	fmt.Println("PaymentProcessor.ProcessPayment:", amount)
	return nil
}

func (s StripeProcessor) RefundPayment(transactionID string, amount float64) error {
	fmt.Println("PaymentProcessor.RefundPayment:", transactionID, amount)
	return nil
}

func (s StripeProcessor) CapturePayment(transactionID string) error {
	fmt.Println("PaymentProcessor.CapturePayment:", transactionID)
	return nil
}

func (s StripeProcessor) VoidPayment(transactionID string) error {
	fmt.Println("PaymentProcessor.VoidPayment:", transactionID)
	return nil
}

// PayPalProcessor implements PaymentProcessor.
type PayPalProcessor struct{}

func (p PayPalProcessor) ProcessPayment(amount float64) error {
	fmt.Println("PaymentProcessor.ProcessPayment:", amount)
	return nil
}

func (p PayPalProcessor) RefundPayment(transactionID string, amount float64) error {
	fmt.Println("PaymentProcessor.RefundPayment:", transactionID, amount)
	return nil
}

func (p PayPalProcessor) CapturePayment(transactionID string) error {
	fmt.Println("PayPalProcessor.CapturePayment not supported")
	return nil
}

func (p PayPalProcessor) VoidPayment(transactionID string) error {
	fmt.Println("PayPalProcessor.VoidPayment not supported")
	return nil
}

func IspI() {
	stripe := StripeProcessor{}
	payPal := PayPalProcessor{}

	stripe.ProcessPayment(100.00)
	stripe.RefundPayment("123", 50.00)
	stripe.CapturePayment("123")
	stripe.VoidPayment("123")

	payPal.ProcessPayment(50.00)
	payPal.RefundPayment("456", 20.00)
}
