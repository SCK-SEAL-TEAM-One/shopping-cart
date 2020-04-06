package tests

import (
	"store-service/internal/product"

	"github.com/stretchr/testify/mock"
)

type mockProductRepository struct {
	mock.Mock
}

func (repo *mockProductRepository) GetProducts(keyword string) (product.ProductResult, error) {
	argument := repo.Called(keyword)
	return argument.Get(0).(product.ProductResult), argument.Error(1)
}

func (repo *mockProductRepository) GetProductByID(ID int) (product.ProductDetail, error) {
	argument := repo.Called(ID)
	return argument.Get(0).(product.ProductDetail), argument.Error(1)
}

func (repo *mockProductRepository) UpdateStock(productID, quantity int) error {
	argument := repo.Called(productID, quantity)
	return argument.Error(0)
}
