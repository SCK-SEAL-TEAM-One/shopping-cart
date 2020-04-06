package payment_test

import (
	"store-service/internal/order"
	"store-service/internal/payment"
	"store-service/internal/product"

	"github.com/stretchr/testify/mock"
)

type mockBankGateway struct {
	mock.Mock
}

func (gateway *mockBankGateway) Payment(paymentDetail payment.PaymentDetail) (string, error) {
	argument := gateway.Called(paymentDetail)
	return argument.String(0), argument.Error(1)
}

type mockShippingGateway struct {
	mock.Mock
}

func (gateway *mockShippingGateway) ShipByKerry(shippingInfo order.ShippingInfo) (string, error) {
	argument := gateway.Called(shippingInfo)
	return argument.String(0), argument.Error(1)
}

type mockOrderRepository struct {
	mock.Mock
}

func (order *mockOrderRepository) GetOrderByShippingMethodByOrderID(orderID int) (string, error) {
	argument := order.Called(orderID)
	return argument.String(0), argument.Error(1)
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

func (order *mockOrderRepository) UpdateOrder(orderID int, transactionID string) error {
	argument := order.Called(orderID, transactionID)
	return argument.Error(0)
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

func (repository *mockOrderRepository) GetOrderProduct(orderID int) ([]order.OrderProduct, error) {
	argument := repository.Called(orderID)
	return argument.Get(0).([]order.OrderProduct), argument.Error(1)
}

type mockShippingRepository struct {
	mock.Mock
}

func (repository *mockShippingRepository) GetShippingByOrderID(orderID int) (order.ShippingInfo, error) {
	argument := repository.Called(orderID)
	return argument.Get(0).(order.ShippingInfo), argument.Error(1)
}
