package api

import (
	"log"
	"net/http"

	"store-service/internal/order"

	"github.com/gin-gonic/gin"
)

type StoreAPI struct {
	OrderService order.OrderInterface
}

type OrderConfirmation struct {
	OrderID    int     `json:"order_id"`
	TotalPrice float64 `json:"total_price"`
}

func (api StoreAPI) SubmitOrderHandler(context *gin.Context) {
	var request order.SubmitedOrder
	if err := context.BindJSON(&request); err != nil {
		context.String(http.StatusBadRequest, err.Error())
		log.Printf("bad request %s", err.Error())
		return
	}

	order := api.OrderService.CreateOrder(request)

	context.JSON(http.StatusOK, OrderConfirmation{
		OrderID:    order.OrderID,
		TotalPrice: order.TotalPrice,
	})
}
