package order

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_GetTotalAmount_Input_SubmittedOrder__Should_Be__(t *testing.T) {
	expectedTotalAmount := 14.95
	var productList = append([]OrderProduct{}, OrderProduct{
		ProductID: 2,
		Quantity:  1,
	})
	submittedOrder := SubmitedOrder{
		Cart:                 productList,
		ShippingMethod:       1,
		ShippingAddress:      "405/35 ถ.มหิดล",
		ShippingSubDistrict:  "ท่าศาลา",
		ShippingDistrict:     "เมือง",
		ShippingProvince:     "เชียงใหม่",
		ShippingZipCode:      "50000",
		RecipientName:        "ณัฐญา ชุติบุตร",
		RecipientPhoneNumber: "0970804292",
	}
	orderService := OrderService{}

	actualTotalAmount := orderService.GetTotalAmount(submittedOrder)

	assert.Equal(t, expectedTotalAmount, actualTotalAmount)
}
