package product

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/go-redis/redis/v7"
	"github.com/jmoiron/sqlx"
)

type ProductRepository interface {
	GetProducts(keyword string, limit string, offset string) (ProductResult, error)
	GetProductByID(ID int) (ProductDetail, error)
	UpdateStock(productID, quantity int) error
}

type ProductRepositoryMySQL struct {
	DBConnection *sqlx.DB
}

func (repository ProductRepositoryMySQL) GetProducts(keyword string, limit string, offset string) (ProductResult, error) {
	var products []Product
	var total int
	if keyword == "" {
		err := repository.DBConnection.Select(&products, "SELECT id,product_name,product_price,image_url FROM products LIMIT ? OFFSET ?", limit, offset)
		if err == nil {
			err = repository.DBConnection.Get(&total, "SELECT count(*) FROM products")
		}

		return ProductResult{
			Total:    total,
			Products: products,
		}, err
	}
	err := repository.DBConnection.Select(&products, "SELECT id,product_name,product_price,image_url FROM products WHERE product_name LIKE ? LIMIT ? OFFSET ?", "%"+keyword+"%", limit, offset)
	if err == nil {
		err = repository.DBConnection.Get(&total, "SELECT count(*) FROM products WHERE product_name LIKE ?", "%"+keyword+"%")
	}

	return ProductResult{
		Total:    total,
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

type ProductRepositoryMySQLWithCache struct {
	RedisConnection *redis.Client
	DBConnection    *sqlx.DB
}

func (repository ProductRepositoryMySQLWithCache) GetProducts(keyword string, limit string, offset string) (ProductResult, error) {
	var products []Product
	var total int

	value, err := repository.RedisConnection.Get(fmt.Sprintf("keyword-%s-%d-%d", keyword, limit, offset)).Result()
	log.Printf("keyword %s value %s error %s", keyword, value, err)
	if err == nil && value != "" {
		var productResult ProductResult
		err = json.Unmarshal([]byte(value), &productResult)
		return ProductResult{
			Total:    productResult.Total,
			Products: productResult.Products,
		}, err
	}

	if keyword == "" {
		err := repository.DBConnection.Select(&products, "SELECT id,product_name,product_price,image_url FROM products LIMIT ? OFFSET ?", limit, offset)
		if err == nil {
			err = repository.DBConnection.Get(&total, "SELECT count(*) FROM products")
		}

		if err == nil {
			data, _ := json.Marshal(ProductResult{
				Total:    total,
				Products: products,
			})
			err = repository.RedisConnection.Set(fmt.Sprintf("keyword-%s-%d-%d", keyword, limit, offset), string(data), time.Hour).Err()
			log.Print("set cache", err)
		}
		log.Print("after query", err)
		return ProductResult{
			Total:    total,
			Products: products,
		}, err
	}
	err = repository.DBConnection.Select(&products, "SELECT id,product_name,product_price,image_url FROM products WHERE product_name LIKE ? LIMIT ? OFFSET ?", "%"+keyword+"%", limit, offset)
	if err == nil {
		err = repository.DBConnection.Get(&total, "SELECT count(*) FROM products WHERE product_name LIKE ?", "%"+keyword+"%")
	}

	if err == nil {
		data, _ := json.Marshal(ProductResult{
			Total:    total,
			Products: products,
		})
		err = repository.RedisConnection.Set(fmt.Sprintf("keyword-%s-%d-%d", keyword, limit, offset), string(data), time.Hour).Err()
		log.Print("set cache", err)
	}
	log.Print("after query", err)
	return ProductResult{
		Total:    total,
		Products: products,
	}, err
}

func (repository ProductRepositoryMySQLWithCache) GetProductByID(ID int) (ProductDetail, error) {
	result := ProductDetail{}

	value, err := repository.RedisConnection.Get(fmt.Sprintf("id-%d", ID)).Result()
	log.Printf("id %d value %s error %s", ID, value, err)
	if err == nil && value != "" {
		err = json.Unmarshal([]byte(value), &result)
		return result, err
	}

	err = repository.DBConnection.Get(&result, "SELECT id,product_name,product_price,quantity,image_url,product_brand FROM products WHERE id=?", ID)
	if err == nil {
		data, _ := json.Marshal(result)
		err = repository.RedisConnection.Set(fmt.Sprintf("id-%d", ID), string(data), time.Hour).Err()
		log.Print("set cache", err)
	}
	log.Print("after query", err)
	return result, err
}

func (productRepository ProductRepositoryMySQLWithCache) UpdateStock(productID, quantity int) error {
	_, err := productRepository.DBConnection.Exec(`UPDATE products SET quantity = quantity-? WHERE id=?`, quantity, productID)
	return err
}
