package order

type OrderInterface interface {
	CreateOrder(totalPrice float64) (int, error)
	CreatedOrderProduct(orderID, productID, quality int, productPrice float64) error
	CreatedShipping(orderID int, shippingInfo ShippingInfo) error
}
