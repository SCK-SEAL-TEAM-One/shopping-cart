package payment

import "github.com/jmoiron/sqlx"

type PaymentRepository interface {
	UpdateStock(productID, quantity int) error
}

type PaymentRepositoryMySQL struct {
	DBConnection *sqlx.DB
}
