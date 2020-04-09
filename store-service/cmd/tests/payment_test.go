package tests

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http/httptest"
	"store-service/cmd/api"
	"store-service/internal/payment"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func Test_ConfirmPaymentHandler_Input_PaymentInformation_Should_Be_ResponsePayment(t *testing.T) {
	expected := `{"notify_message":"วันเวลาที่ชำระเงิน 1/3/2563 13:30:00 หมายเลขคำสั่งซื้อ 8004359103 คุณสามารถติดตามสินค้าผ่านช่องทาง Kerry หมายเลข Tracking 1785261900"}
`
	paymentInformation := payment.PaymentInformation{
		OrderID:      1337620837,
		PaymentType:  "credit",
		Type:         "visa",
		CardNumber:   "4719700591590995",
		CVV:          "752",
		ExpiredMonth: 7,
		ExpiredYear:  20,
		CardName:     "Karnwat Wongudom",
		TotalPrice:   102.00,
	}

	requestJSON, _ := json.Marshal(paymentInformation)
	request := httptest.NewRequest("POST", "/api/v1/confirmPayment", bytes.NewBuffer(requestJSON))
	write := httptest.NewRecorder()

	orderID := 1337620837
	mockPaymentService := new(mockPaymentService)
	mockPaymentService.On("ConfirmPayment", orderID, payment.PaymentDetail{
		CardNumber:   "4719700591590995",
		CVV:          "752",
		ExpiredMonth: 7,
		ExpiredYear:  20,
		CardName:     "Karnwat Wongudom",
		TotalPrice:   102.00,
	}).Return("วันเวลาที่ชำระเงิน 1/3/2563 13:30:00 หมายเลขคำสั่งซื้อ 8004359103 คุณสามารถติดตามสินค้าผ่านช่องทาง Kerry หมายเลข Tracking 1785261900", nil)

	paymentAPI := api.PaymentAPI{
		PaymentService: mockPaymentService,
	}

	mockRoute := gin.Default()
	mockRoute.POST("/api/v1/confirmPayment", paymentAPI.ConfirmPaymentHandler)
	mockRoute.ServeHTTP(write, request)
	response := write.Result()
	actual, err := ioutil.ReadAll(response.Body)

	assert.Equal(t, nil, err)
	assert.Equal(t, expected, string(actual))
}
