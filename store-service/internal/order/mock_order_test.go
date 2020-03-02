package order_test

import (
	"store-service/internal/order"
	"store-service/internal/product"

	"github.com/stretchr/testify/mock"
)

type mockOrder struct {
	mock.Mock
}

func (order *mockOrder) CreateOrder(totalPrice float64) (int, error) {
	argument := order.Called(totalPrice)
	return argument.Int(0), argument.Error(1)
}

func (order *mockOrder) CreatedOrderProduct(orderID int, listProduct []order.OrderProduct) error {
	argument := order.Called(orderID, listProduct)
	return argument.Error(0)
}

func (order *mockOrder) CreatedShipping(orderID int, shippingInfo order.ShippingInfo) error {
	argument := order.Called(orderID, shippingInfo)
	return argument.Error(0)
}

func (order *mockOrder) GetProductByID(id int) product.Product {
	argument := order.Called(id)
	return argument.Get(0).(product.Product)
}
