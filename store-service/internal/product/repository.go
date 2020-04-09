package product

import (
	"github.com/jmoiron/sqlx"
)

type ProductRepository interface {
	GetProducts(keyword string) (ProductResult, error)
	GetProductByID(ID int) (ProductDetail, error)
	UpdateStock(productID, quantity int) error
}

type ProductRepositoryMySQL struct {
	DBConnection *sqlx.DB
}

func (repository ProductRepositoryMySQL) GetProducts(keyword string) (ProductResult, error) {
	var products []Product
	if keyword == "" {
		err := repository.DBConnection.Select(&products, "SELECT id,product_name,product_price,image_url FROM products")
		return ProductResult{
			Total:    len(products),
			Products: products,
		}, err
	}
	err := repository.DBConnection.Select(&products, "SELECT id,product_name,product_price,image_url FROM products WHERE produt_name = ?%", keyword)
	return ProductResult{
		Total:    len(products),
		Products: products,
	}, err
}

func (productRepository ProductRepositoryMySQL) GetProductByID(ID int) (ProductDetail, error) {
	result := ProductDetail{}
	err := productRepository.DBConnection.Get(&result, "SELECT id,product_name,product_price,quantity,image_url,product_brand FROM products WHERE id=?", ID)
	return result, err
}

func (productRepository ProductRepositoryMySQL) UpdateStock(productID, quantity int) error {
	_, err := productRepository.DBConnection.Exec(`UPDATE products SET quantity = quantity-? WHERE id=?`, quantity, productID)
	return err
}
