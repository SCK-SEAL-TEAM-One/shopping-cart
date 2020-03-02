package order

import "github.com/jmoiron/sqlx"

type OrderRepositoryMySQL struct {
	DBConnection *sqlx.DB
}

func (orderRepository OrderRepositoryMySQL) CreateShipping(orderID int, submittedOrder SubmitedOrder) (int, error) {
	result := orderRepository.
		DBConnection.
		MustExec(`INSERT INTO shipping (order_id, address, sub_district, district, province, zip_code, recipient, phone_number) VALUES (?, ?, ?, ?, ?, ?, ?, ?)`,
			orderID,
			submittedOrder.ShippingAddress,
			submittedOrder.ShippingSubDistrict,
			submittedOrder.ShippingDistrict,
			submittedOrder.ShippingProvince,
			submittedOrder.ShippingZipCode,
			submittedOrder.RecipientName,
			submittedOrder.RecipientPhoneNumber,
		)
	id, err := result.LastInsertId()
	return int(id), err
}
