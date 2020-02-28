package api

import (
	"log"
	"net/http"

	"store-service/internal/order"

	"github.com/gin-gonic/gin"
)

type StoreAPI struct {
	OrderDB order.OrderInterface
}
type SubmmitedOrder struct {
	Cart                 []order.OrderProduct `json:"cart"`
	ShippingMethod       int                  `json:"shipping_method"`
	ShippingAddress      string               `json:"shipping_address"`
	ShippingSubDistrict  string               `json:"shipping_sub_disterict"`
	ShippingDistrict     string               `json:"shipping_district"`
	ShippingProvince     string               `json:"shipping_province"`
	ShippingZipCode      string               `json:"shipping_zip_code"`
	RecipientName        string               `json:"recipient_name"`
	RecipientPhoneNumber string               `json:"recipient_phone_number"`
}

type OrderConfirmation struct {
	OrderID    int     `json:"order_id"`
	TotalPrice float64 `json:"total_price"`
}

func (api StoreAPI) SubmitOrderHandler(context *gin.Context) {
	var request SubmmitedOrder
	if err := context.BindJSON(&request); err != nil {
		context.String(http.StatusBadRequest, err.Error())
		log.Printf("bad request %s", err.Error())
		return
	}

	totalPrice := GetTotalProductPrice(request.Cart)
	orderID, err := api.OrderDB.CreateOrder(totalPrice)
	if err != nil {
		context.String(http.StatusInternalServerError, err.Error())
		log.Printf("GetTotalProductPrice internal error %s", err.Error())
		return
	}

	if err := api.OrderDB.CreatedOrderProduct(orderID, request.Cart); err != nil {
		context.String(http.StatusInternalServerError, err.Error())
		log.Printf("CreatedOrderProduct internal error %s", err.Error())
		return
	}

	shippingInfo := order.ShippingInfo{
		ShippingMethod:       request.ShippingMethod,
		ShippingAddress:      request.ShippingAddress,
		ShippingSubDistrict:  request.ShippingSubDistrict,
		ShippingDistrict:     request.ShippingDistrict,
		ShippingProvince:     request.ShippingProvince,
		ShippingZipCode:      request.ShippingZipCode,
		RecipientName:        request.RecipientName,
		RecipientPhoneNumber: request.RecipientPhoneNumber,
	}
	if err := api.OrderDB.CreatedShipping(orderID, shippingInfo); err != nil {
		context.String(http.StatusInternalServerError, err.Error())
		log.Printf("CreatedShipping internal error %s", err.Error())
		return
	}

	context.JSON(http.StatusOK, OrderConfirmation{
		OrderID:    orderID,
		TotalPrice: totalPrice,
	})
}

func GetTotalProductPrice(orderProduct []order.OrderProduct) float64 {
	return 12.95
}
