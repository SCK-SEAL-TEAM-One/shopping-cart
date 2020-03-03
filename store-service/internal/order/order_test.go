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
		TotalPrice: 14.95,
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

	mockProductRepository := new(mockProductRepository)
	mockProductRepository.On("GetProductByID", 2).Return(product.ProductDetail{
		ID:       2,
		Name:     "43 Piece dinner Set",
		Price:    12.95,
		Quantity: 1,
		Brand:    "Coolkidz",
		Image:    "43_Piece_Dinner_Set.jpg",
	}, nil)

	mockOrderRepository := new(mockOrderRepository)
	orderID := 8004359103
	productID := 2
	quantity := 1
	totalPrice := 14.95
	productPrice := 12.95

	mockOrderRepository.On("CreateOrder", totalPrice).Return(orderID, nil)

	mockOrderRepository.On("CreateOrderProduct", orderID, productID, quantity, productPrice).Return(nil)

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
	mockOrderRepository.On("CreateShipping", orderID, shippingInfo).Return(1, nil)

	orderService := order.OrderService{
		ProductRepository: mockProductRepository,
		OrderRepository:   mockOrderRepository,
	}

	actual := orderService.CreateOrder(submittedOrder)

	assert.Equal(t, expected, actual)
}

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
	mockProductRepository.On("GetProductByID", 2).Return(product.ProductDetail{
		ID:       2,
		Name:     "43 Piece dinner Set",
		Price:    12.95,
		Quantity: 1,
		Brand:    "Coolkidz",
	}, nil)

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
	mockProductRepository.On("GetProductByID", 2).Return(product.ProductDetail{
		ID:       2,
		Name:     "43 Piece dinner Set",
		Price:    12.95,
		Quantity: 1,
		Brand:    "Coolkidz",
	}, nil)

	orderService := order.OrderService{
		ProductRepository: mockProductRepository,
	}

	actualTotalAmount := orderService.GetTotalAmount(submittedOrder)

	assert.Equal(t, expectedTotalAmount, actualTotalAmount)
}
