package order_test

import "github.com/stretchr/testify/mock"

type mockProductRepository struct {
	mock.Mock
}

func (repository ProductRepositoryMySQL) GetProductByID(id int) product.Product {
	argument := order.Called(id)
	return argument.Get(0).(product.Product)
}
