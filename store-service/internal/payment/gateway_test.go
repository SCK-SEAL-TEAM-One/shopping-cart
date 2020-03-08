//+build integration

package payment_test

import (
	"store-service/internal/payment"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Payment_Input_PaymentDetail_CardNumber_4719700591590995_Should_Be_TransactionID_TOY202002021525(t *testing.T) {
	expectedTransactionID := "TOY202002021525"

	paymentDetail := payment.PaymentDetail{
		CardNumber:   "4719700591590995",
		CVV:          "752",
		ExpiredMonth: 7,
		ExpiredYear:  20,
		CardName:     "Karnwat Wongudom",
		TotalPrice:   104.95,
		MerchantID:   154124000,
	}

	gateway := payment.BankGateway{
		BankEndpoint: "http://localhost:8882",
	}
	actualTransactionID, err := gateway.Payment(paymentDetail)

	assert.Equal(t, expectedTransactionID, actualTransactionID)
	assert.Equal(t, nil, err)

}
