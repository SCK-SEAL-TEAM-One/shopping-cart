package order

import "github.com/jmoiron/sqlx"

type OrderRepositoryMySQL struct {
	DBConnection *sqlx.DB
}

func (orderRepository OrderRepositoryMySQL) CreateShipping(orderID int, submittedOrder SubmitedOrder) (int, error) {
	result := orderRepository.DBConnection.MustExec(`INSERT INTO shipping (order_id, address, sub_district, districtd, province, zip_code, recipiant, phone_number) 
		VALUE ($1, $2, $3, $4, $5, $6, $7, $8)`,
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
