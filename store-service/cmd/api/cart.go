package api

import (
	"log"
	"net/http"

	"store-service/internal/cart"

	"github.com/gin-gonic/gin"
)

type CartAPI struct {
	CartService cart.CartService
}

// type OrderConfirmation struct {
// 	OrderID    int     `json:"order_id"`
// 	TotalPrice float64 `json:"total_price"`
// }

func (api CartAPI) AddCartHandler(context *gin.Context) {
	var request cart.SubmitedCart
	if err := context.BindJSON(&request); err != nil {
		context.String(http.StatusBadRequest, err.Error())
		log.Printf("bad request %s", err.Error())
		return
	}

	act, err := api.CartService.AddCart(request)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	// context.JSON(http.StatusOK, cart)
	context.JSON(http.StatusOK, gin.H{
		"status": act,
	})
}
