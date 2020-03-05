//+build integration

package shipping_test

import (
	"store-service/internal/order"
	"store-service/internal/shipping"
	"testing"

	_ "github.com/go-sql-driver/mysql"

	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
)

func Test_ShippingRepository(t *testing.T) {
	connection, err := sqlx.Connect("mysql", "sealteam:sckshuhari@(localhost:3306)/toy")
	if err != nil {
		t.Fatalf("cannot tearup data err %s", err)
	}
	repository := shipping.ShippingRepositoryMySQL{
		DBConnection: connection,
	}

	t.Run("GetShippingByOrderID_Input_OrderID_0_Should_Be_ShippingInfo", func(t *testing.T) {
		expectedShippingInfo := order.ShippingInfo{
			ShippingMethod:       "Kerry",
			ShippingAddress:      "405/37 ถ.มหิดล",
			ShippingSubDistrict:  "ท่าศาลา",
			ShippingDistrict:     "เมือง",
			ShippingProvince:     "เชียงใหม่",
			ShippingZipCode:      "50000",
			RecipientName:        "ณัฐญา ชุติบุตร",
			RecipientPhoneNumber: "0970809292",
		}
		orderID := 1

		actualShippingInfo, err := repository.GetShippingByOrderID(orderID)

		assert.Equal(t, expectedShippingInfo, actualShippingInfo)
		assert.Equal(t, nil, err)
	})
}
