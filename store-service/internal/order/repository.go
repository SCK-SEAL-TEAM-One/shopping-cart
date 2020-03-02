package order

import "github.com/jmoiron/sqlx"

type OrderRepository interface {
	CreateOrder(totalPrice float64) (int, error)
	CreateOrderProduct(orderID, productID, quantity int, productPrice float64) error
	CreateShipping(orderID int, shippingInfo ShippingInfo) (int, error)
}

type OrderRepositoryMySQL struct {
	DBConnection *sqlx.DB
}

func (orderRepository OrderRepositoryMySQL) CreateShipping(orderID int, shippingInfo ShippingInfo) (int, error) {
	result := orderRepository.
		DBConnection.
		MustExec(`INSERT INTO shipping (order_id, address, sub_district, district, province, zip_code, recipient, phone_number) VALUES (?, ?, ?, ?, ?, ?, ?, ?)`,
			orderID,
			shippingInfo.ShippingAddress,
			shippingInfo.ShippingSubDistrict,
			shippingInfo.ShippingDistrict,
			shippingInfo.ShippingProvince,
			shippingInfo.ShippingZipCode,
			shippingInfo.RecipientName,
			shippingInfo.RecipientPhoneNumber,
		)
	id, err := result.LastInsertId()
	return int(id), err
}

func (orderRepository OrderRepositoryMySQL) CreateOrder(totalPrice float64) (int, error) {
	tx := orderRepository.DBConnection.MustBegin()
	sqlResult := tx.MustExec("INSERT INTO orders (total_price) VALUE (?)", totalPrice)
	insertedId, err := sqlResult.LastInsertId()
	return int(insertedId), err
}

func (orderRepository OrderRepositoryMySQL) CreateOrderProduct(orderID int, productID, quantity int, productPrice float64) error {
	tx := orderRepository.DBConnection.MustBegin()
	sqlResult := tx.MustExec("INSERT INTO order_product (order_id, product_id) VALUE (?,?)", orderID, productID)
	_, err := sqlResult.RowsAffected()
	return err
}
