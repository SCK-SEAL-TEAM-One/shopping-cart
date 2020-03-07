package payment

import (
	"store-service/internal/order"
	"store-service/internal/product"
	"store-service/internal/shipping"
	"time"
)

type PaymentInterface interface {
	ConfirmPayment(orderID int, paymentdetail PaymentDetail) string
}

type BankGatewayInterface interface {
	Payment(paymentDetail PaymentDetail) (string, error)
}

type ShippingGatewayInterface interface {
	ShipByKerry(shippingInfo order.ShippingInfo) (string, error)
}

type PaymentService struct {
	BankGateway        BankGatewayInterface
	ShippingGateway    ShippingGatewayInterface
	OrderRepository    order.OrderRepository
	ProductRepository  product.ProductRepository
	ShippingRepository shipping.ShippingRepository
	Time               time.Time
}

func (service PaymentService) ConfirmPayment(orderID int, paymentdetail PaymentDetail) string {
	transactionId, err := service.BankGateway.Payment(paymentdetail)
	if err != nil {
		return ""
	}
	orderProductList, err := service.OrderRepository.GetOrderProduct(orderID)
	if err != nil {
		return ""
	}
	for _, orderProduct := range orderProductList {
		err = service.ProductRepository.UpdateStock(orderProduct.ProductID, orderProduct.Quantity)
		if err != nil {
			return ""
		}
	}
	shippingInfo, err := service.ShippingRepository.GetShippingByOrderID(orderID)
	if err != nil {
		return ""
	}
	trackingID, err := service.ShippingGateway.ShipByKerry(shippingInfo)
	if err != nil {
		return ""
	}
	err = service.OrderRepository.UpdateOrder(orderID, transactionId)
	if err != nil {
		return ""
	}

	return order.SendNotification(orderID, trackingID, service.Time, shippingInfo.ShippingMethod)
}
