package cart_test

import (
	"database/sql"
	"store-service/internal/cart"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_GetCart_Should_be_Length(t *testing.T) {
	expected := 1
	uid := 1
	res := []cart.CartDetail{
		{
			ID:        1,
			UserID:    1,
			ProductID: 2,
			Quantity:  1,
			Name:      "43 Piece dinner Set",
			Price:     12.95,
			Image:     "/43_Piece_dinner_Set.png",
			Stock:     10,
			Brand:     "CoolKidz",
		},
	}
	mockCartRepository := new(mockCartRepository)
	mockCartRepository.On("GetCartDetail", uid).Return(res, nil)

	cartService := cart.CartService{
		CartRepository: mockCartRepository,
	}
	actual, err := cartService.GetCart(uid)

	assert.Len(t, actual, expected)
	assert.Equal(t, nil, err)
}

func Test_GetCart_Should_be_Empty(t *testing.T) {
	expected := 0
	uid := 1
	res := []cart.CartDetail{}
	mockCartRepository := new(mockCartRepository)
	mockCartRepository.On("GetCartDetail", uid).Return(res, nil)

	cartService := cart.CartService{
		CartRepository: mockCartRepository,
	}
	actual, err := cartService.GetCart(uid)

	assert.Len(t, actual, expected)
	assert.Equal(t, nil, err)
}

func Test_AddCart_Input_Submitted_Cart_Should_be_Added(t *testing.T) {
	expected := "added"
	submitedCart := cart.SubmitedCart{
		ProductID: 1,
		Quantity:  1,
	}
	uid := 1
	mockCartRepository := new(mockCartRepository)
	mockCartRepository.On("GetCartByProductID", uid, submitedCart.ProductID).Return(cart.Cart{}, sql.ErrNoRows)
	mockCartRepository.On("CreateCart", uid, submitedCart.ProductID, submitedCart.Quantity).Return(1, nil)

	cartService := cart.CartService{
		CartRepository: mockCartRepository,
	}
	actual, err := cartService.AddCart(uid, submitedCart)

	assert.Equal(t, expected, actual)
	assert.Equal(t, nil, err)
}

func Test_AddCart_Input_Submitted_Cart_Should_be_Updated(t *testing.T) {
	expected := "updated"
	submitedCart := cart.SubmitedCart{
		ProductID: 1,
		Quantity:  1,
	}
	uid := 1
	mockCartRepository := new(mockCartRepository)
	mockCartRepository.On("GetCartByProductID", uid, submitedCart.ProductID).Return(cart.Cart{}, nil)
	mockCartRepository.On("UpdateCart", uid, submitedCart.ProductID, submitedCart.Quantity).Return(nil)

	cartService := cart.CartService{
		CartRepository: mockCartRepository,
	}
	actual, err := cartService.AddCart(uid, submitedCart)

	assert.Equal(t, expected, actual)
	assert.Equal(t, nil, err)
}

func Test_UpdateCart_Input_Submitted_Cart_Should_be_Updated(t *testing.T) {
	expected := "updated"
	submitedCart := cart.SubmitedCart{
		ProductID: 1,
		Quantity:  1,
	}
	uid := 1
	mockCartRepository := new(mockCartRepository)
	mockCartRepository.On("UpdateCart", uid, submitedCart.ProductID, submitedCart.Quantity).Return(nil)

	cartService := cart.CartService{
		CartRepository: mockCartRepository,
	}
	actual, err := cartService.UpdateCart(uid, submitedCart)

	assert.Equal(t, expected, actual)
	assert.Equal(t, nil, err)
}

func Test_UpdateCart_Input_Submitted_Cart_Should_be_Deleted(t *testing.T) {
	expected := "deleted"
	submitedCart := cart.SubmitedCart{
		ProductID: 1,
		Quantity:  0,
	}
	uid := 1
	mockCartRepository := new(mockCartRepository)
	mockCartRepository.On("DeleteCart", uid, submitedCart.ProductID).Return(nil)

	cartService := cart.CartService{
		CartRepository: mockCartRepository,
	}
	actual, err := cartService.UpdateCart(uid, submitedCart)

	assert.Equal(t, expected, actual)
	assert.Equal(t, nil, err)
}
