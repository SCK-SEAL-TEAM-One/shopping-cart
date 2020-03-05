package main

import (
	"log"
	"store-service/cmd/api"
	"store-service/internal/healthcheck"
	"store-service/internal/order"
	"store-service/internal/payment"
	"store-service/internal/product"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func main() {
	connection, err := sqlx.Connect("mysql", "sealteam:sckshuhari@(store-database:3306)/toy")
	if err != nil {
		log.Fatalln("cannot connect to database", err)
	}
	productRepository := product.ProductRepositoryMySQL{
		DBConnection: connection,
	}
	orderRepository := order.OrderRepositoryMySQL{
		DBConnection: connection,
	}
	orderService := order.OrderService{
		ProductRepository: &productRepository,
		OrderRepository:   &orderRepository,
	}
	paymentService := payment.PaymentService{}
	storeAPI := api.StoreAPI{
		OrderService: &orderService,
	}
	paymentAPI := api.PaymentAPI{
		PaymentService: &paymentService,
	}

	route := gin.Default()
	route.POST("/api/v1/order", storeAPI.SubmitOrderHandler)
	route.POST("/api/v1/confirmPayment", paymentAPI.ConfirmPaymentHandler)

	route.GET("/api/v1/health", func(context *gin.Context) {
		user, err := healthcheck.GetUserNameFromDB(connection)
		if err != nil {
			context.JSON(500, gin.H{
				"error": err.Error(),
			})
			return
		}
		context.JSON(200, gin.H{
			"message": user,
		})
	})
	log.Fatal(route.Run(":8000"))
}
