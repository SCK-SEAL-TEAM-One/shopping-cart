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
