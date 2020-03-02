package order

import (
	"log"
	"store-service/internal/product"
)

type OrderService struct {
	ProductRepository product.ProductRepository
	OrderRepository   OrderInterface
}

func (service OrderService) CreateOrder(submitedOrder SubmitedOrder) Order {
	var orderProduct OrderProduct

	productID, err := service.ProductRepository.GetProductByID(orderProduct.ProductID)
	if err != nil {
		log.Printf("GetProductByID internal error %s", err.Error())
		return Order{}
	}

	totalAmount := GetTotalAmount(submitedOrder)

	orderID, err := service.OrderRepository.CreateOrder(totalAmount)
	if err != nil {
		log.Printf("OrderRepository.CreateOrder internal error %s", err.Error())
		return Order{}
	}

	shippingInfo := ShippingInfo{
		ShippingMethod:       submitedOrder.ShippingMethod,
		ShippingAddress:      submitedOrder.ShippingAddress,
		ShippingSubDistrict:  submitedOrder.ShippingSubDistrict,
		ShippingDistrict:     submitedOrder.ShippingDistrict,
		ShippingProvince:     submitedOrder.ShippingProvince,
		ShippingZipCode:      submitedOrder.ShippingZipCode,
		RecipientName:        submitedOrder.RecipientName,
		RecipientPhoneNumber: submitedOrder.RecipientPhoneNumber,
	}
	err = service.OrderRepository.CreatedShipping(orderID, shippingInfo)
	if err != nil {
		log.Printf("OrderRepository.CreatedShipping internal error %s", err.Error())
		return Order{}
	}

	for _, selectedProduct := range submitedOrder.Cart {
		err = service.OrderRepository.CreatedOrderProduct(orderID, selectedProduct.ProductID, selectedProduct.Quantity, selectedProduct.ProductPrice)
		if err != nil {
			log.Printf("OrderRepository.CreatedOrderProduct internal error %s", err.Error())
			return Order{}
		}
	}
	return Order{}
}
