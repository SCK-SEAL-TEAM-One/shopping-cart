package product

import "github.com/jmoiron/sqlx"

type ProductRepository interface {
	GetProductByID(ID int) (Product, error)
}

type ProductRepositoryMySQL struct {
	DBConnection *sqlx.DB
}

func (product ProductRepositoryMySQL) GetProductByID(ID int) (Product, error) {
	return Product{}, nil
}
