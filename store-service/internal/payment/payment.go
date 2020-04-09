package payment

import (
	"log"
	"store-service/internal/order"
	"store-service/internal/product"
	"store-service/internal/shipping"
	"time"
)

type PaymentInterface interface {
	ConfirmPayment(orderID int, paymentdetail PaymentDetail) (string, error)
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
	Time               func() time.Time
}

func (service PaymentService) ConfirmPayment(orderID int, paymentdetail PaymentDetail) (string, error) {
	transactionId, err := service.BankGateway.Payment(paymentdetail)
	if err != nil {
		log.Printf("BankGateway.Payment internal error %s", err.Error())
		return "", err
	}
	orderProductList, err := service.OrderRepository.GetOrderProduct(orderID)
	if err != nil {
		log.Printf("OrderRepository.GetOrderProduct internal error %s", err.Error())
		return "", err
	}
	for _, orderProduct := range orderProductList {
		err = service.ProductRepository.UpdateStock(orderProduct.ProductID, orderProduct.Quantity)
		if err != nil {
			log.Printf("ProductRepository.UpdateStock internal error %s", err.Error())
			return "", err
		}
	}
	shippingInfo, err := service.ShippingRepository.GetShippingByOrderID(orderID)
	if err != nil {
		log.Printf("ShippingRepository.GetShippingByOrderID internal error %s", err.Error())
		return "", err
	}
	trackingID, err := service.ShippingGateway.ShipByKerry(shippingInfo)
	if err != nil {
		log.Printf("ShippingGateway.ShipByKerry internal error %s", err.Error())
		return "", err
	}
	err = service.OrderRepository.UpdateOrder(orderID, transactionId)
	if err != nil {
		log.Printf("OrderRepository.UpdateOrder internal error %s", err.Error())
		return "", err
	}

	return order.SendNotification(orderID, trackingID, service.Time(), shippingInfo.ShippingMethod), nil
}
