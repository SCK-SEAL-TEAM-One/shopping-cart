package order

type OrderInterface interface {
	CreateOrder(totalPrice float64) (int, error)
	CreatedOrderProduct(orderID int, listProduct []OrderProduct) error
	CreatedShipping(orderID int, shippingInfo ShippingInfo) error
}

type OrderService struct {
}

func (service OrderService) GetTotalProductPrice(order SubmitedOrder) float64 {
	return 12.95
}

func (service OrderService) GetTotalAmount(order SubmitedOrder) float64 {
	return service.GetTotalProductPrice(order) + order.GetShippingFee()
}
