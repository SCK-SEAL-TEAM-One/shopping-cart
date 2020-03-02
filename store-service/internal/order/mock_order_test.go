package order_test

import (
	"store-service/internal/product"

	"github.com/stretchr/testify/mock"
)

type mockProductRepository struct {
	mock.Mock
}

func (repository mockProductRepository) GetProductByID(id int) product.Product {
	argument := repository.Called(id)
	return argument.Get(0).(product.Product)
}
