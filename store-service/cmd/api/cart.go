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

func (api CartAPI) GetCartHandler(context *gin.Context) {
	uid := 1
	cart, err := api.CartService.GetCart(uid)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	context.JSON(http.StatusOK, cart)
}

func (api CartAPI) AddCartHandler(context *gin.Context) {
	var request cart.SubmitedCart
	if err := context.BindJSON(&request); err != nil {
		context.String(http.StatusBadRequest, err.Error())
		log.Printf("bad request %s", err.Error())
		return
	}

	uid := 1
	act, err := api.CartService.AddCart(uid, request)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"status": act,
	})
}

func (api CartAPI) UpdateCartHandler(context *gin.Context) {
	var request cart.SubmitedCart
	if err := context.BindJSON(&request); err != nil {
		context.String(http.StatusBadRequest, err.Error())
		log.Printf("bad request %s", err.Error())
		return
	}

	uid := 1
	act, err := api.CartService.UpdateCart(uid, request)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"status": act,
	})
}
