package cart_test

import (
	"store-service/internal/cart"

	"github.com/stretchr/testify/mock"
)

type mockCartRepository struct {
	mock.Mock
}

func (repo *mockCartRepository) GetCartDetail(userID int) ([]cart.CartDetail, error) {
	argument := repo.Called(userID)
	return argument.Get(0).([]cart.CartDetail), argument.Error(1)
}

func (repo *mockCartRepository) GetCartByProductID(userID int, productID int) (cart.Cart, error) {
	argument := repo.Called(userID, productID)
	return argument.Get(0).(cart.Cart), argument.Error(1)
}

func (repo *mockCartRepository) CreateCart(userID int, productID int, quantity int) (int, error) {
	argument := repo.Called(userID, productID, quantity)
	return argument.Int(0), argument.Error(1)
}

func (repo *mockCartRepository) UpdateCart(userID int, productID int, quantity int) error {
	argument := repo.Called(userID, productID, quantity)
	return argument.Error(0)
}

func (repo *mockCartRepository) DeleteCart(userID int, productID int) error {
	argument := repo.Called(userID, productID)
	return argument.Error(0)
}
