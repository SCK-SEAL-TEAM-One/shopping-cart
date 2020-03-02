// +build integration

package order_test

import (
	"store-service/internal/order"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_CreateShipping_Input_OrderID_8004359103_Should_Be_ShippingID_1_No_Error(t *testing.T) {
	expectShippingID := 1
	productList := []order.OrderProduct{
		{
			ProductID: 2,
			Quantity:  1,
		},
	}
	submittedOrder := order.SubmitedOrder{
		Cart:                 productList,
		ShippingMethod:       1,
		ShippingAddress:      "405/35 ถ.มหิดล",
		ShippingSubDistrict:  "ท่าศาลา",
		ShippingDistrict:     "เมือง",
		ShippingProvince:     "เชียงใหม่",
		ShippingZipCode:      "50000",
		RecipientName:        "ณัฐญา ชุติบุตร",
		RecipientPhoneNumber: "0970804292",
	}
	orderID := 8004359103
	orderRepository := order.OrderRepositoryMySQL{}

	actualShippingID, err := orderRepository.CreateShipping(orderID, submittedOrder)
	assert.Equal(t, expectShippingID, actualShippingID)
	assert.Equal(t, err, nil)

}
