package order_test

import (
	"store-service/internal/order"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_GetShippingFee_Input_SubmitedOrder_ShippingMethod_1_Should_Be_Fee_2(t *testing.T) {
	expectedFee := 2.00
	submitOrder := order.SubmitedOrder{
		ShippingMethod: "Kerry",
	}
	actualFee := submitOrder.GetShippingFee()

	assert.Equal(t, expectedFee, actualFee)
}
