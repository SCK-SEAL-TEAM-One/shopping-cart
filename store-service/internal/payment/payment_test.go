package payment_test

import (
	"errors"
	"store-service/internal/order"
	"store-service/internal/payment"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_ConfirmPayment_Input_OrderID_8004359103_And_PaymentDetail_Should_Be_NotificationMessage(t *testing.T) {
	expectedMessage := "วันเวลาที่ชำระเงิน 1/3/2020 13:30:00 หมายเลขคำสั่งซื้อ 8004359103 คุณสามารถติดตามสินค้าผ่านช่องทาง Kerry หมายเลข 1785261900"

	orderId := 8004359103
	paymentDetail := payment.PaymentDetail{
		CardNumber:   "4719700591590995",
		CVV:          "752",
		ExpiredMonth: 7,
		ExpiredYear:  20,
		CardName:     "Karnwat Wongudom",
		TotalPrice:   104.95,
		MerchantID:   154124000,
	}
	shippingInfo := order.ShippingInfo{
		ShippingMethod:       "Kerry",
		ShippingAddress:      "405/35 ถ.มหิดล",
		ShippingSubDistrict:  "ท่าศาลา",
		ShippingDistrict:     "เมือง",
		ShippingProvince:     "เชียงใหม่",
		ShippingZipCode:      "50000",
		RecipientName:        "ณัฐญา ชุติบุตร",
		RecipientPhoneNumber: "0970804292",
	}

	mockBankGateway := new(mockBankGateway)
	mockBankGateway.On("Payment", paymentDetail).Return("TOY202002021525", nil)

	mockShippingGateway := new(mockShippingGateway)
	mockShippingGateway.On("ShipByKerry", shippingInfo).Return("1785261900", nil)

	mockOrderRepository := new(mockOrderRepository)
	mockOrderRepository.On("GetOrderProduct", orderId).Return([]order.OrderProduct{
		{
			ProductID: 2,
			Quantity:  1,
		},
	}, nil)
	mockOrderRepository.On("UpdateOrder", orderId, "TOY202002021525").Return(nil)

	mockProductRepository := new(mockProductRepository)
	mockProductRepository.On("UpdateStock", 2, 1).Return(nil)

	mockShippingRepository := new(mockShippingRepository)
	mockShippingRepository.On("GetShippingByOrderID", orderId).Return(shippingInfo, nil)

	fixedTime, _ := time.Parse("2/1/2006T15:04:05", "1/3/2020T13:30:00")

	paymentService := payment.PaymentService{
		BankGateway:        mockBankGateway,
		ShippingGateway:    mockShippingGateway,
		OrderRepository:    mockOrderRepository,
		ProductRepository:  mockProductRepository,
		ShippingRepository: mockShippingRepository,
		Time: func() time.Time {
			return fixedTime
		},
	}

	actualMessage, err := paymentService.ConfirmPayment(orderId, paymentDetail)
	assert.Equal(t, expectedMessage, actualMessage)
	assert.Equal(t, nil, err)
}

func Test_ConfirmPayment_Input_OrderID_8004359103_And_PaymentDetail_Should_Be_BankGateway_Return_Error(t *testing.T) {
	expectedMessage := ""

	orderId := 8004359103
	paymentDetail := payment.PaymentDetail{
		CardNumber:   "4719700591590995",
		CVV:          "752",
		ExpiredMonth: 7,
		ExpiredYear:  20,
		CardName:     "Karnwat Wongudom",
		TotalPrice:   104.95,
		MerchantID:   154124000,
	}

	mockBankGateway := new(mockBankGateway)
	mockBankGateway.On("Payment", paymentDetail).Return("", errors.New("BankGateway Error"))

	paymentService := payment.PaymentService{
		BankGateway: mockBankGateway,
	}

	actualMessage, err := paymentService.ConfirmPayment(orderId, paymentDetail)
	assert.Equal(t, expectedMessage, actualMessage)
	assert.Equal(t, errors.New("BankGateway Error"), err)
}
