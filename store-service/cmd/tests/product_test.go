package tests

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"store-service/cmd/api"
	"store-service/internal/product"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func Test_ProductSearchHandler_Should_Be_ProductResult(t *testing.T) {
	expected := `{"total":1,"products":[{"id":1,"product_name":"43 Piece Dinner Set","product_price":10,"product_image":"/43-piece-dinner-set.png"}]}
`

	request := httptest.NewRequest("GET", "/api/v1/product", nil)
	write := httptest.NewRecorder()

	mockProductRepository := new(mockProductRepository)
	mockProductRepository.On("GetProducts", "").Return(product.ProductResult{
		Total: 1,
		Products: []product.Product{
			{
				ID:    1,
				Name:  "43 Piece Dinner Set",
				Price: 10.00,
				Image: "/43-piece-dinner-set.png",
			},
		},
	}, nil)

	api := api.ProductAPI{
		ProductRepository: mockProductRepository,
	}

	mockRoute := gin.Default()
	mockRoute.GET("/api/v1/product", api.SearchHandler)
	mockRoute.ServeHTTP(write, request)
	response := write.Result()
	actual, err := ioutil.ReadAll(response.Body)

	assert.Equal(t, nil, err)
	assert.Equal(t, http.StatusOK, response.StatusCode)
	assert.Equal(t, expected, string(actual))
}

func Test_ProductGetProductHandler_Should_Be_ProductResult(t *testing.T) {
	expected := `{"id":1,"product_name":"43 Piece Dinner Set","product_price":10,"product_image":"/43-piece-dinner-set.png","quantity":2,"product_brand":"CoolKidz"}
`

	request := httptest.NewRequest("GET", "/api/v1/product/1", nil)
	write := httptest.NewRecorder()

	mockProductRepository := new(mockProductRepository)
	mockProductRepository.On("GetProductByID", 1).Return(product.ProductDetail{
		ID:       1,
		Name:     "43 Piece Dinner Set",
		Price:    10.00,
		Image:    "/43-piece-dinner-set.png",
		Quantity: 2,
		Brand:    "CoolKidz",
	}, nil)

	api := api.ProductAPI{
		ProductRepository: mockProductRepository,
	}

	mockRoute := gin.Default()
	mockRoute.GET("/api/v1/product/:id", api.GetProductHandler)
	mockRoute.ServeHTTP(write, request)
	response := write.Result()
	actual, err := ioutil.ReadAll(response.Body)

	assert.Equal(t, nil, err)
	assert.Equal(t, http.StatusOK, response.StatusCode)
	assert.Equal(t, expected, string(actual))
}
