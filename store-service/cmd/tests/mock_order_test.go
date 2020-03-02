package tests

import (
	"store-service/internal/order"

	"github.com/stretchr/testify/mock"
)

type mockOrderDB struct {
	mock.Mock
}

func (orderService *mockOrderDB) CreateOrder(submitedOrder order.SubmitedOrder) order.Order {
	argument := orderService.Called(submitedOrder)
	return argument.Get(0).(order.Order)
}
