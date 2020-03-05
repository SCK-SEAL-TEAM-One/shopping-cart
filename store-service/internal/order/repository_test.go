//+build integration

package order_test

import (
	"fmt"
	"store-service/internal/order"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
)

func Test_OrderRepository(t *testing.T) {
	connection, err := sqlx.Connect("mysql", "sealteam:sckshuhari@(localhost:3306)/toy")
	if err != nil {
		t.Fatalf("cannot tearup data err %s", err)
	}
	repository := order.OrderRepositoryMySQL{
		DBConnection: connection,
	}

	t.Run("CreateOrder_Input_TotalPrice_14_dot_95_ShippingMethod_Kerry_Should_Be_OrderID_No_Error", func(t *testing.T) {
		totalPrice := 14.95
		shippingMethod := "Kerry"

		actualId, err := repository.CreateOrder(totalPrice, shippingMethod)

		assert.Equal(t, nil, err)
		assert.NotEmpty(t, actualId)
	})

	t.Run("CreateOrderProduct_Input_OrderID_2_And_ProductID_2_Should_Be_No_Error", func(t *testing.T) {
		orderId := 2
		productId := 2
		quantity := 1
		productPrice := 12.95
		err := repository.CreateOrderProduct(orderId, productId, quantity, productPrice)

		assert.Equal(t, nil, err)
	})

	t.Run("CreateShipping_Input_OrderID_8004359103_Should_Be_ShippingID_1_No_Error", func(t *testing.T) {
		submittedOrder := order.ShippingInfo{
			ShippingMethod:       "Kerry",
			ShippingAddress:      "405/35 ถ.มหิดล",
			ShippingSubDistrict:  "ท่าศาลา",
			ShippingDistrict:     "เมือง",
			ShippingProvince:     "เชียงใหม่",
			ShippingZipCode:      "50000",
			RecipientName:        "ณัฐญา ชุติบุตร",
			RecipientPhoneNumber: "0970804292",
		}
		orderID := 8004359103
		orderRepository := order.OrderRepositoryMySQL{
			DBConnection: connection,
		}

		actualShippingID, err := orderRepository.CreateShipping(orderID, submittedOrder)
		assert.NotEmpty(t, actualShippingID)
		assert.Equal(t, err, nil)
	})

	t.Run("UpdateOrder_Input_ToyID_TOY202002021525_OrderID_8004359103_Should_No_Error", func(t *testing.T) {
		transactionID := "TOY202002021525"
		orderID := 8004359104

		err := repository.UpdateOrder(orderID, transactionID)

		assert.Equal(t, nil, err)
	})

	t.Run("UpdateOrder_Input_ToyID_TOY202002021525_OrderID_11111111119_Should_Get_Error_No_Row_Affected", func(t *testing.T) {
		expectedError := fmt.Errorf("no any row affected , update not completed")
		transactionID := "TOY202002021525"
		orderID := 11111111119

		err := repository.UpdateOrder(orderID, transactionID)

		assert.Equal(t, expectedError, err)
	})

	t.Run("GetOrderProduct_Input_OrderID_1_Should_Be_OrderProducts", func(t *testing.T) {
		expectedOrderProducts := []order.OrderProduct{
			{
				ProductID: 2,
				Quantity:  10,
			},
			{
				ProductID: 1,
				Quantity:  10,
			},
		}

		orderID := 1

		actualOrderProducts, err := repository.GetOrderProduct(orderID)

		assert.Equal(t, expectedOrderProducts, actualOrderProducts)
		assert.Equal(t, nil, err)
	})
}
