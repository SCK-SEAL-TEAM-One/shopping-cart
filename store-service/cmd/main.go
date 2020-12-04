package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"store-service/cmd/api"
	"store-service/internal/healthcheck"
	"store-service/internal/order"
	"store-service/internal/payment"
	"store-service/internal/product"
	"store-service/internal/shipping"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v7"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"

	"go.elastic.co/apm/module/apmgin"
)

func main() {
	var testMode, useCache bool
	if os.Getenv("TEST_MODE") != "" {
		testMode = true
	}
	if os.Getenv("CACHE_ON") != "" {
		useCache = true
	}

	bankGatewayEndpoint := "bank-gateway:8882"
	shippingGatewayEndpoint := "shipping-gateway:8882"

	if os.Getenv("BANK_GATEWAY") != "" {
		bankGatewayEndpoint = os.Getenv("BANK_GATEWAY")
	}
	if os.Getenv("SHIPPING_GATEWAY") != "" {
		shippingGatewayEndpoint = os.Getenv("SHIPPING_GATEWAY")
	}

	dbConnecton := "sealteam:sckshuhari@(store-database:3306)/toy"
	if os.Getenv("DBCONNECTION") != "" {
		dbConnecton = os.Getenv("DBCONNECTION")
	}
	connection, err := sqlx.Connect("mysql", dbConnecton)
	if err != nil {
		log.Fatalln("cannot connect to database", err)
	}
	var productRepository product.ProductRepository

	if useCache {
		rdb := redis.NewClient(&redis.Options{
			Addr:     "store-cache:6379", // use default Addr
			Password: "",                 // no password set
			DB:       0,                  // use default DB
		})
		_, err := rdb.Ping().Result()
		if err != nil {
			log.Fatalln("cannot connect to cache", err)
		}
		log.Println("redis connected")
		productRepository = &product.ProductRepositoryMySQLWithCache{
			DBConnection:    connection,
			RedisConnection: rdb,
		}
	} else {
		productRepository = &product.ProductRepositoryMySQL{
			DBConnection: connection,
		}
	}

	orderRepository := order.OrderRepositoryMySQL{
		DBConnection: connection,
	}
	shippingRepository := shipping.ShippingRepositoryMySQL{
		DBConnection: connection,
	}
	orderService := order.OrderService{
		ProductRepository: productRepository,
		OrderRepository:   &orderRepository,
	}
	bankGateway := payment.BankGateway{
		BankEndpoint: "http://" + bankGatewayEndpoint,
	}
	shippingGateway := shipping.ShippingGateway{
		KerryEndpoint: "http://" + shippingGatewayEndpoint,
	}
	paymentService := payment.PaymentService{
		BankGateway:        &bankGateway,
		ShippingGateway:    &shippingGateway,
		OrderRepository:    &orderRepository,
		ProductRepository:  productRepository,
		ShippingRepository: &shippingRepository,
		Time:               time.Now,
	}
	storeAPI := api.StoreAPI{
		OrderService: &orderService,
	}
	paymentAPI := api.PaymentAPI{
		PaymentService: &paymentService,
	}
	productAPI := api.ProductAPI{
		ProductRepository: productRepository,
	}

	route := gin.Default()
	route.Use(apmgin.Middleware(route))
	route.GET("/api/v1/product", productAPI.SearchHandler)
	route.GET("/api/v1/product/:id", productAPI.GetProductHandler)
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
	if testMode {
		log.Println("Test mode: ", testMode)
		route.GET("/mockTime/:time", func(context *gin.Context) {
			fixedTime, err := time.Parse("02012006T15:04:05", context.Param("time"))
			if err != nil {
				context.JSON(http.StatusBadRequest, gin.H{
					"status": "fail",
					"masess": err,
				})
				return
			}
			now := func() time.Time {
				return fixedTime
			}
			paymentService.Time = now
			context.JSON(http.StatusOK, gin.H{
				"status":  "ok",
				"fixTime": fixedTime,
			})
		})
	}

	log.Fatal(route.Run(":8000"))
}

func GetUserNameFromDB(connection *sqlx.DB) User {
	user := User{}
	err := connection.Get(&user, "SELECT id,name FROM user WHERE ID=1")
	if err != nil {
		fmt.Printf("Get user name from tearup get error : %s", err.Error())
		return User{}
	}
	return user
}

type User struct {
	ID   int    `db:"id"`
	Name string `db:"name"`
}
