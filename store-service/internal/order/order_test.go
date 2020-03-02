package order_test

import (
	"store-service/internal/order"
	"store-service/internal/product"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_CreateOrder_Input_Submitted_Order_Should_be_OrderID_8004359103_TotalPrice_100_Dot_00(t *testing.T) {
	expected := order.Order{
		OrderID:    8004359103,
		TotalPrice: 100.00,
	}

	submittedOrder := order.SubmitedOrder{
		Cart: []order.OrderProduct{
			{
				ProductID: 2,
				Quantity:  1,
			},
		},
		ShippingMethod:       1,
		ShippingAddress:      "405/37 ถ.มหิดล",
		ShippingSubDistrict:  "ท่าศาลา",
		ShippingDistrict:     "เมือง",
		ShippingProvince:     "เชียงใหม่",
		ShippingZipCode:      "50000",
		RecipientName:        "ณัฐญา ชุติบุตร",
		RecipientPhoneNumber: "0970809292",
	}

	mockOrderDB := new(mockOrder)

	mockOrderDB.On("GetProductByID", 2).Return(product.Product{
		ID:    2,
		Name:  "43 Piece Dinner Set",
		Price: 12.95,
		Image: "43PieceDinnerSet.jpg",
	})

	orderID := 8004359103
	mockOrderDB.On("CreateOrder", 12.95).Return(orderID, nil)
	listProduct := []order.OrderProduct{
		{
			ProductID: 2,
			Quantity:  1,
		},
	}

	mockOrderDB.On("CreatedOrderProduct", orderID, listProduct).Return(nil)

	shippingInfo := order.ShippingInfo{
		ShippingMethod:       1,
		ShippingAddress:      "405/37 ถ.มหิดล",
		ShippingSubDistrict:  "ท่าศาลา",
		ShippingDistrict:     "เมือง",
		ShippingProvince:     "เชียงใหม่",
		ShippingZipCode:      "50000",
		RecipientName:        "ณัฐญา ชุติบุตร",
		RecipientPhoneNumber: "0970809292",
	}
	mockOrderDB.On("CreatedShipping", orderID, shippingInfo).Return(nil)

	actual := CreateOrder(submittedOrder)

	assert.Equal(t, expected, actual)
}
