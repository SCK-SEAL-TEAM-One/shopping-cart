package order_test

import (
	"store-service/internal/order"
	"store-service/internal/product"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_GetTotalProductPrice_Input_SummitedOrder_Cart_ProductID_2_Quantity_1_Should_Be_TotalProductPrice_12_dot_95(t *testing.T) {
	expectedTotalProductPrice := 12.95
	submitOrder := order.SubmitedOrder{
		ShippingMethod: 1,
		Cart: []order.OrderProduct{
			{
				ProductID: 2,
				Quantity:  1,
			},
		},
	}

	mockProductRepository := new(mockProductRepository)
	mockProductRepository.On("GetProductByID", 2).Return(product.Product{
		ID:       2,
		Name:     "43 Piece dinner Set",
		Price:    12.95,
		Quantity: 1,
		Brand:    "Coolkidz",
	})

	orderService := order.OrderService{
		ProductRepository: mockProductRepository,
	}
	actualTotalPrice := orderService.GetTotalProductPrice(submitOrder)

	assert.Equal(t, expectedTotalProductPrice, actualTotalPrice)
}

func Test_GetTotalAmount_Input_SubmittedOrder_ProductID_2_Quantity_1_Should_Be_TotalPrice_14_dot_95(t *testing.T) {
	expectedTotalAmount := 14.95
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
	mockProductRepository := new(mockProductRepository)
	mockProductRepository.On("GetProductByID", 2).Return(product.Product{
		ID:       2,
		Name:     "43 Piece dinner Set",
		Price:    12.95,
		Quantity: 1,
		Brand:    "Coolkidz",
	})
	orderService := order.OrderService{
		ProductRepository: mockProductRepository,
	}

	actualTotalAmount := orderService.GetTotalAmount(submittedOrder)

	assert.Equal(t, expectedTotalAmount, actualTotalAmount)
}
