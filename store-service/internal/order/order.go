package order

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
