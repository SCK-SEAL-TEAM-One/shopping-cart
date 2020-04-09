package api

import (
	"log"
	"net/http"
	"store-service/internal/payment"

	"github.com/gin-gonic/gin"
)

type PaymentAPI struct {
	PaymentService payment.PaymentInterface
}

func (api PaymentAPI) ConfirmPaymentHandler(context *gin.Context) {
	var request payment.PaymentInformation
	if err := context.BindJSON(&request); err != nil {
		context.String(http.StatusBadRequest, err.Error())
		log.Printf("bad request %s", err.Error())
		return
	}

	paymentDetail := payment.NewShippingInfo(request)
	payment, err := api.PaymentService.ConfirmPayment(request.OrderID, paymentDetail)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"notify_message": payment,
	})
}
