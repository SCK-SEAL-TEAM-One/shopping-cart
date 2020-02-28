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

func Test_SubmitOrderHandler_Input_Order_One_Piece_Should_Be_Order_ID_1337620837_And_12_Dot_95(t *testing.T) {
	expected := `{"order_id":1337620837,"total_price":12.95}
`
	submittedOrder := api.SubmmitedOrder{
		Cart: []order.OrderProduct{
			{
				ProductID: 2,
				Quantity:  1,
			},
		},
		ShippingMethod:       1,
		ShippingAddress:      "405/37 ถ.มหิดล",
		ShippingSubDistrict:  "ท่าศาลา",
		ShippingDistrict:     "เมือง",
		ShippingProvince:     "เชียงใหม่",
		ShippingZipCode:      "50000",
		RecipientName:        "ณัฐญา ชุติบุตร",
		RecipientPhoneNumber: "0970809292",
	}

	requestJson, _ := json.Marshal(submittedOrder)
	request := httptest.NewRequest("POST", "/api/v1/order", bytes.NewBuffer(requestJson))
	write := httptest.NewRecorder()

	mockOrderDB := new(mockOrderDB)
	orderID := 1337620837
	mockOrderDB.On("CreateOrder", 12.95).Return(orderID, nil)
	listProduct := []order.OrderProduct{
		{
			ProductID: 2,
			Quantity:  1,
		},
	}
	mockOrderDB.On("CreatedOrderProduct", orderID, listProduct).Return(nil)
	shippingInfo := order.ShippingInfo{
		ShippingMethod:       1,
		ShippingAddress:      "405/37 ถ.มหิดล",
		ShippingSubDistrict:  "ท่าศาลา",
		ShippingDistrict:     "เมือง",
		ShippingProvince:     "เชียงใหม่",
		ShippingZipCode:      "50000",
		RecipientName:        "ณัฐญา ชุติบุตร",
		RecipientPhoneNumber: "0970809292",
	}
	mockOrderDB.On("CreatedShipping", orderID, shippingInfo).Return(nil)

	storeAPI := api.StoreAPI{
		OrderDB: mockOrderDB,
	}
	mockRoute := gin.Default()
	mockRoute.POST("/api/v1/order", storeAPI.SubmitOrderHandler)
	mockRoute.ServeHTTP(write, request)
	response := write.Result()
	actual, err := ioutil.ReadAll(response.Body)

	assert.Equal(t, nil, err)
	assert.Equal(t, expected, string(actual))
}
