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

	confirm := payment.PaymentDetail{
		CardNumber:   request.CardNumber,
		CVV:          request.CVV,
		ExpiredMonth: request.ExpiredMonth,
		ExpiredYear:  request.ExpiredYear,
		CardName:     request.CardName,
		TotalPrice:   request.TotalPrice,
	}

	payment := api.PaymentService.ConfirmPayment(request.OrderID, confirm)

	context.JSON(http.StatusOK, gin.H{
		"notify_message": payment,
	})
}
