package order

import "store-service/internal/product"

type OrderInterface interface {
	CreateOrder(totalPrice float64) (int, error)
	CreatedOrderProduct(orderID int, listProduct []OrderProduct) error
	CreatedShipping(orderID int, shippingInfo ShippingInfo) error
}

type OrderService struct {
	ProductRepository ProductRepository
}

type ProductRepository interface {
	GetProductByID(id int) product.Product
}

func (orderService OrderService) GetTotalProductPrice(submitedOrder SubmitedOrder) float64 {
	totalProductPrice := 0.00
	for _, cartItem := range submitedOrder.Cart {
		product := orderService.ProductRepository.GetProductByID(cartItem.ProductID)
		totalProductPrice += product.Price * float64(cartItem.Quantity)
	}
	return totalProductPrice
}

func (orderService OrderService) GetTotalAmount(order SubmitedOrder) float64 {
	return orderService.GetTotalProductPrice(order) + order.GetShippingFee()
}
