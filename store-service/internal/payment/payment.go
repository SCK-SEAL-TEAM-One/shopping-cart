package payment

type PaymentInterface interface {
	ConfirmPayment(paymentdetail PaymentDetail) string
}
