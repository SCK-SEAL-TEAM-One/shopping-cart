package tests

import (
	"store-service/internal/order"

	"github.com/stretchr/testify/mock"
)

type mockOrderDB struct {
	mock.Mock
}

func (order *mockOrderDB) CreateOrder(totalPrice float64) (int, error) {
	argument := order.Called(totalPrice)
	return argument.Int(0), argument.Error(1)
}

func (order *mockOrderDB) CreatedOrderProduct(orderID, productID, quantity int, productPrice float64) error {
	argument := order.Called(orderID, productID, quantity, productPrice)
	return argument.Error(0)
}

func (order *mockOrderDB) CreatedShipping(orderID int, shippingInfo order.ShippingInfo) error {
	argument := order.Called(orderID, shippingInfo)
	return argument.Error(0)
}
