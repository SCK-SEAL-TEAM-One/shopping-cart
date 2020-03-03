package payment

import "github.com/jmoiron/sqlx"

type PaymentRepository interface {
	UpdateStock(productID, quantity int) error
}

type PaymentRepositoryMySQL struct {
	DBConnection *sqlx.DB
}

func (paymentRepository PaymentRepositoryMySQL) UpdateStock(productID, quantity int) error {
	_, err := paymentRepository.DBConnection.Exec(`UPDATE products SET quantity = quantity-? WHERE id=?`, quantity, productID)
	return err
}
