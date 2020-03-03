package tests

import (
	"store-service/internal/order"

	"github.com/stretchr/testify/mock"
)

type mockOrderService struct {
	mock.Mock
}

func (orderService *mockOrderService) CreateOrder(submitedOrder order.SubmitedOrder) order.Order {
	argument := orderService.Called(submitedOrder)
	return argument.Get(0).(order.Order)
}
