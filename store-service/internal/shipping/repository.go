package shipping

import (
	"store-service/internal/order"

	"github.com/jmoiron/sqlx"
)

type ShippingRepository interface {
	GetShippingByOrderID(orderID int) (order.ShippingInfo, error)
}

type ShippingRepositoryMySQL struct {
	DBConnection *sqlx.DB
}

func (repository ShippingRepositoryMySQL) GetShippingByOrderID(orderID int) (order.ShippingInfo, error) {
	var shippingInfo order.ShippingInfo
	err := repository.DBConnection.Get(&shippingInfo, "SELECT orders.shipping_method as method, address, sub_district, district, province, zip_code, recipient, phone_number FROM shipping INNER JOIN orders ON shipping.order_id = orders.id WHERE order_id = ?", orderID)
	return shippingInfo, err
}
