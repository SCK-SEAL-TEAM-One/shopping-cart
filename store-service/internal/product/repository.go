package product

import "github.com/jmoiron/sqlx"

type ProductRepository interface {
	GetProductByID(ID int) (ProductDetail, error)
}

type ProductRepositoryMySQL struct {
	DBConnection *sqlx.DB
}

func (productRepository ProductRepositoryMySQL) GetProductByID(ID int) (ProductDetail, error) {
	result := ProductDetail{}
	err := productRepository.DBConnection.Get(&result, "SELECT id,product_name,product_price,quantity,image_url,product_brand FROM products WHERE id=?", ID)
	return result, err
}
