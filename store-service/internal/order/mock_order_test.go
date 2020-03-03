package order_test

import (
	"store-service/internal/order"
	"store-service/internal/product"

	"github.com/stretchr/testify/mock"
)

type mockOrderRepository struct {
	mock.Mock
}

func (order *mockOrderRepository) CreateOrder(totalPrice float64, shippingMethod string) (int, error) {
	argument := order.Called(totalPrice, shippingMethod)
	return argument.Int(0), argument.Error(1)
}

func (order *mockOrderRepository) CreateOrderProduct(orderID, productID, quantity int, productPrice float64) error {
	argument := order.Called(orderID, productID, quantity, productPrice)
	return argument.Error(0)
}

func (order *mockOrderRepository) CreateShipping(orderID int, shippingInfo order.ShippingInfo) (int, error) {
	argument := order.Called(orderID, shippingInfo)
	return argument.Int(0), argument.Error(1)
}

type mockProductRepository struct {
	mock.Mock
}

func (repository *mockProductRepository) GetProductByID(id int) (product.ProductDetail, error) {
	argument := repository.Called(id)
	return argument.Get(0).(product.ProductDetail), argument.Error(1)
}
