package cart

import (
	"database/sql"
	"log"
)

type CartService struct {
	CartRepository CartRepository
}

type CartInterface interface {
	AddCart(submitedCart SubmitedCart) (Cart, error)
}

func (cartService CartService) AddCart(uid int, submitedCart SubmitedCart) (string, error) {
	cart, err := cartService.CartRepository.GetCartByProductID(uid, submitedCart.ProductID)
	act := "updated"
	if err == sql.ErrNoRows {
		act = "added"
		cartService.CartRepository.CreateCart(uid, submitedCart.ProductID, submitedCart.Quantity)
		return act, nil
	}
	err = cartService.CartRepository.UpdateCart(uid, submitedCart.ProductID, submitedCart.Quantity+cart.Quantity)
	if err != nil {
		log.Printf("CartRepository.UpdateCart internal error %s", err.Error())
		return "", err
	}
	return act, nil
}

func (cartService CartService) UpdateCart(uid int, submitedCart SubmitedCart) (string, error) {
	act := "updated"
	if submitedCart.Quantity == 0 {
		act = "deleted"
		err := cartService.CartRepository.DeleteCart(uid, submitedCart.ProductID)
		if err != nil {
			log.Printf("CartRepository.DeleteCart internal error %s", err.Error())
			return "", err
		}
	} else {
		err := cartService.CartRepository.UpdateCart(uid, submitedCart.ProductID, submitedCart.Quantity)
		if err != nil {
			log.Printf("CartRepository.UpdateCart internal error %s", err.Error())
			return "", err
		}
	}
	return act, nil
}

// func (orderService OrderService) CreateOrder(submitedOrder SubmitedOrder) Order {
// 	totalPrice := orderService.GetTotalAmount(submitedOrder)

// 	orderID, err := orderService.OrderRepository.CreateOrder(totalPrice, submitedOrder.GetShippingMethodProvider())
// 	if err != nil {
// 		log.Printf("OrderRepository.CreateOrder internal error %s", err.Error())
// 		return Order{}
// 	}

// 	shippingInfo := ShippingInfo{
// 		ShippingMethod:       submitedOrder.ShippingMethod,
// 		ShippingAddress:      submitedOrder.ShippingAddress,
// 		ShippingSubDistrict:  submitedOrder.ShippingSubDistrict,
// 		ShippingDistrict:     submitedOrder.ShippingDistrict,
// 		ShippingProvince:     submitedOrder.ShippingProvince,
// 		ShippingZipCode:      submitedOrder.ShippingZipCode,
// 		RecipientName:        submitedOrder.RecipientName,
// 		RecipientPhoneNumber: submitedOrder.RecipientPhoneNumber,
// 	}
// 	_, err = orderService.OrderRepository.CreateShipping(orderID, shippingInfo)
// 	if err != nil {
// 		log.Printf("OrderRepository.CreateShipping internal error %s", err.Error())
// 		return Order{}
// 	}

// 	for _, selectedProduct := range submitedOrder.Cart {
// 		product, err := orderService.ProductRepository.GetProductByID(selectedProduct.ProductID)
// 		err = orderService.OrderRepository.CreateOrderProduct(orderID, selectedProduct.ProductID, selectedProduct.Quantity, product.Price)
// 		if err != nil {
// 			log.Printf("OrderRepository.CreateOrderProduct internal error %s", err.Error())
// 			return Order{}
// 		}
// 	}
// 	return Order{
// 		OrderID:    orderID,
// 		TotalPrice: totalPrice,
// 	}
// }

// func (orderService OrderService) GetTotalProductPrice(submitedOrder SubmitedOrder) float64 {
// 	totalProductPrice := 0.00
// 	for _, cartItem := range submitedOrder.Cart {
// 		product, _ := orderService.ProductRepository.GetProductByID(cartItem.ProductID)
// 		totalProductPrice += product.Price * float64(cartItem.Quantity)
// 	}
// 	return totalProductPrice
// }

// func (orderService OrderService) GetTotalAmount(order SubmitedOrder) float64 {
// 	return orderService.GetTotalProductPrice(order) + order.GetShippingFee()
// }

// func SendNotification(orderID int, trackingNumber string, dateTime time.Time, shippingMethod string) string {
// 	return fmt.Sprintf("วันเวลาที่ชำระเงิน %s หมายเลขคำสั่งซื้อ %d คุณสามารถติดตามสินค้าผ่านช่องทาง %s หมายเลข %s", dateTime.Format("2/1/2006 15:04:05"), orderID, shippingMethod, trackingNumber)
// }
