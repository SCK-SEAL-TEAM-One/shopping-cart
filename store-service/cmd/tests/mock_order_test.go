package tests

import (
	"store-service/internal/order"
	"time"

	"github.com/stretchr/testify/mock"
)

type mockOrderService struct {
	mock.Mock
}

func (orderService *mockOrderService) CreateOrder(submitedOrder order.SubmitedOrder) order.Order {
	argument := orderService.Called(submitedOrder)
	return argument.Get(0).(order.Order)
}

func (orderService *mockOrderService) SendNotification(orderID, trackingID int, dateTime time.Time, shippingMethod string) string {
	argument := orderService.Called(orderID, trackingID, dateTime, shippingMethod)
	return argument.String(0)
}
