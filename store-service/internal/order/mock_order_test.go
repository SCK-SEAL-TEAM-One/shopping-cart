package order_test

import (
	"store-service/internal/order"
	"store-service/internal/product"

	"github.com/stretchr/testify/mock"
)

type mockOrderRepository struct {
	mock.Mock
}

func (repo *mockOrderRepository) GetOrderByShippingMethodByOrderID(orderID int) (string, error) {
	argument := repo.Called(orderID)
	return argument.String(0), argument.Error(1)
}

func (repo *mockOrderRepository) CreateOrder(totalPrice float64, shippingMethod string) (int, error) {
	argument := repo.Called(totalPrice, shippingMethod)
	return argument.Int(0), argument.Error(1)
}

func (repo *mockOrderRepository) CreateOrderProduct(orderID, productID, quantity int, productPrice float64) error {
	argument := repo.Called(orderID, productID, quantity, productPrice)
	return argument.Error(0)
}

func (repo *mockOrderRepository) GetOrderProduct(orderID int) ([]order.OrderProduct, error) {
	argument := repo.Called(orderID)
	return argument.Get(0).([]order.OrderProduct), argument.Error(1)
}

func (repo *mockOrderRepository) CreateShipping(orderID int, shippingInfo order.ShippingInfo) (int, error) {
	argument := repo.Called(orderID, shippingInfo)
	return argument.Int(0), argument.Error(1)
}

func (repo *mockOrderRepository) UpdateOrder(orderID int, transactionID string) error {
	argument := repo.Called(orderID, transactionID)
	return argument.Error(1)
}

type mockProductRepository struct {
	mock.Mock
}

func (repo *mockProductRepository) GetProducts(keyword string) (product.ProductResult, error) {
	argument := repo.Called(keyword)
	return argument.Get(0).(product.ProductResult), argument.Error(1)
}

func (repository *mockProductRepository) GetProductByID(id int) (product.ProductDetail, error) {
	argument := repository.Called(id)
	return argument.Get(0).(product.ProductDetail), argument.Error(1)
}

func (repository *mockProductRepository) UpdateStock(productId, quantity int) error {
	argument := repository.Called(productId, quantity)
	return argument.Error(0)
}
