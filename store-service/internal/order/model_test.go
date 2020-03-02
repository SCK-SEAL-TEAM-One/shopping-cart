package order_test

import (
	"store-service/internal/order"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_GetShippingFee_Input_SubmitOrder_ShippingMethod_1_Should_Be_Fee_2(t *testing.T) {
	expectedFee := 2.00
	submitOrder := order.SubmitedOrder{
		ShippingMethod: 1,
	}
	actualFee := submitOrder.GetShippingFee()

	assert.Equal(t, expectedFee, actualFee)
}
