package tests

import (
	"store-service/internal/payment"

	"github.com/stretchr/testify/mock"
)

type mockPaymentService struct {
	mock.Mock
}

func (paymentService *mockPaymentService) ConfirmPayment(paymentDetail payment.PaymentDetail) string {
	argument := paymentService.Called(paymentDetail)
	return argument.String(0)
}
