package order_test

import (
	"store-service/internal/order"
	"store-service/internal/product"

	"github.com/stretchr/testify/mock"
)

type mockOrderRepository struct {
	mock.Mock
}

func (order *mockOrderRepository) CreateOrder(totalPrice float64) (int, error) {
	argument := order.Called(totalPrice)
	return argument.Int(0), argument.Error(1)
}

func (order *mockOrderRepository) CreatedOrderProduct(orderID, productID, quantity int, productPrice float64) error {
	argument := order.Called(orderID, productID, quantity, productPrice)
	return argument.Error(0)
}

func (order *mockOrderRepository) CreatedShipping(orderID int, shippingInfo order.ShippingInfo) error {
	argument := order.Called(orderID, shippingInfo)
	return argument.Error(0)
}

type mockProductRepository struct {
	mock.Mock
}

func (repository *mockProductRepository) GetProductByID(id int) (product.Product, error) {
	argument := repository.Called(id)
	return argument.Get(0).(product.Product), argument.Error(1)
}
