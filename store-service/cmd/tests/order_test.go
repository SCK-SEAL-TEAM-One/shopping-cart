package tests

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http/httptest"
	"store-service/cmd/api"
	"store-service/internal/order"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func Test_SubmitOrderHandler_Input_Order_One_Piece_Should_Be_Order_ID_1337620837_And_14_Dot_95(t *testing.T) {
	expected := `{"order_id":1337620837,"total_price":14.95}
`

	submittedOrder := order.SubmitedOrder{
		Cart: []order.OrderProduct{
			{
				ProductID: 2,
				Quantity:  1,
			},
		},
		ShippingMethod:       "Kerry",
		ShippingAddress:      "405/37 ถ.มหิดล",
		ShippingSubDistrict:  "ท่าศาลา",
		ShippingDistrict:     "เมือง",
		ShippingProvince:     "เชียงใหม่",
		ShippingZipCode:      "50000",
		RecipientName:        "ณัฐญา ชุติบุตร",
		RecipientPhoneNumber: "0970809292",
	}

	requestJSON, _ := json.Marshal(submittedOrder)
	request := httptest.NewRequest("POST", "/api/v1/order", bytes.NewBuffer(requestJSON))
	write := httptest.NewRecorder()

	mockOrderService := new(mockOrderService)
	orderID := 1337620837
	mockOrderService.On("CreateOrder", order.SubmitedOrder{
		Cart: []order.OrderProduct{
			{
				ProductID: 2,
				Quantity:  1,
			},
		},
		ShippingMethod:       "Kerry",
		ShippingAddress:      "405/37 ถ.มหิดล",
		ShippingSubDistrict:  "ท่าศาลา",
		ShippingDistrict:     "เมือง",
		ShippingProvince:     "เชียงใหม่",
		ShippingZipCode:      "50000",
		RecipientName:        "ณัฐญา ชุติบุตร",
		RecipientPhoneNumber: "0970809292",
	}).Return(order.Order{
		OrderID:    orderID,
		TotalPrice: 14.95,
	}, nil)

	storeAPI := api.StoreAPI{
		OrderService: mockOrderService,
	}

	mockRoute := gin.Default()
	mockRoute.POST("/api/v1/order", storeAPI.SubmitOrderHandler)
	mockRoute.ServeHTTP(write, request)
	response := write.Result()
	actual, err := ioutil.ReadAll(response.Body)

	assert.Equal(t, nil, err)
	assert.Equal(t, expected, string(actual))
}
