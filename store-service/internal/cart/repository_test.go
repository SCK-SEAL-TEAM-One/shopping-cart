//go:build integration
// +build integration

package cart_test

import (
	"fmt"
	"store-service/internal/cart"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
)

func Test_CartRepository(t *testing.T) {
	connection, err := sqlx.Connect("mysql", "sealteam:sckshuhari@(localhost:3306)/toy")
	if err != nil {
		t.Fatalf("cannot tearup data err %s", err)
	}
	repository := cart.CartRepositoryMySQL{
		DBConnection: connection,
	}

	t.Run("CreateCart_Input_ProductID_2_Quantity_1_Should_Be_CartID_No_Error", func(t *testing.T) {
		uid, pid, qty := 1, 2, 1
		actualId, err := repository.CreateCart(uid, pid, qty)

		assert.Equal(t, nil, err)
		assert.NotEmpty(t, actualId)
	})

	t.Run("UpdateCart_Input_ProductID_2_Quantity_2_Should_Be_No_Error", func(t *testing.T) {
		uid, pid, qty := 1, 2, 2

		err := repository.UpdateCart(uid, pid, qty)
		assert.Equal(t, nil, err)
	})

	t.Run("UpdateCart_Input_ProductID_3_Quantity_2_Should_Be_Error", func(t *testing.T) {
		uid, pid, qty := 1, 3, 2
		expectedError := fmt.Errorf("no any row affected , update not completed")

		err := repository.UpdateCart(uid, pid, qty)
		assert.Equal(t, expectedError, err)
	})

	t.Run("GetCartDetail_Input_Exist_UserID_Should_Be_Length_No_Error", func(t *testing.T) {
		uid := 1

		actualCarts, err := repository.GetCartDetail(uid)
		assert.Len(t, actualCarts, 1)
		assert.Equal(t, nil, err)
	})

	t.Run("GetCartDetail_Input_Not_Exist_UserID_Should_Be_No_Length_No_Error", func(t *testing.T) {
		uid := 0

		actualCarts, err := repository.GetCartDetail(uid)
		assert.Len(t, actualCarts, 0)
		assert.Equal(t, nil, err)
	})

	t.Run("GetCartByProductID_Input_ProductID_Should_Be_Cart_No_Error", func(t *testing.T) {
		uid, pid := 1, 2
		expected := cart.Cart{
			ID:        1,
			UserID:    1,
			ProductID: 2,
			Quantity:  2,
		}

		actualCart, err := repository.GetCartByProductID(uid, pid)
		assert.Equal(t, expected, actualCart)
		assert.Equal(t, nil, err)
	})

	t.Run("DeleteCart_Input_ProductID_Should_Be_No_Error", func(t *testing.T) {
		uid, pid := 1, 2

		err := repository.DeleteCart(uid, pid)
		assert.Equal(t, nil, err)
	})
}
