//+build integration

package shipping_test

import (
	"store-service/internal/order"
	"store-service/internal/shipping"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ShipByKerry_Input_ShippingInfo_Should_Be_Tracking_Number_1785261900_And_No_Error(t *testing.T) {
	expectedTrackingNumber := "1785261900"

	shippingInfo := order.ShippingInfo{
		ShippingMethod:       "Kerry",
		ShippingAddress:      "405/35 ถ.มหิดล",
		ShippingSubDistrict:  "ท่าศาลา",
		ShippingDistrict:     "เมือง",
		ShippingProvince:     "เชียงใหม่",
		ShippingZipCode:      "50000",
		RecipientName:        "ณัฐญา ชุติบุตร",
		RecipientPhoneNumber: "0970804292",
	}

	service := shipping.ShippingGateway{
		KerryEndpoint: "http://localhost:8883",
	}
	actualTrackingNumber, err := service.ShipByKerry(shippingInfo)

	assert.Equal(t, nil, err)
	assert.Equal(t, expectedTrackingNumber, actualTrackingNumber)
}
