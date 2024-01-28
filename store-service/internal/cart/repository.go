package cart

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

type CartRepository interface {
	GetCartByProductID(userID int, productID int) (Cart, error)
	CreateCart(userID int, productID int, quantity int) (int, error)
	UpdateCart(userID int, productID int, quantity int) error
	DeleteCart(userID int, productID int) error
}

type CartRepositoryMySQL struct {
	DBConnection *sqlx.DB
}

func (repository CartRepositoryMySQL) GetCartByProductID(userID int, productID int) (Cart, error) {
	var cart Cart
	err := repository.DBConnection.Get(&cart, "SELECT id,user_id,product_id,quantity FROM carts WHERE user_id = ? AND product_id = ? LIMIT 1", userID, productID)
	return cart, err
}

func (repository CartRepositoryMySQL) CreateCart(userID int, productID int, quantity int) (int, error) {
	sqlResult := repository.DBConnection.MustExec("INSERT INTO carts (user_id, product_id, quantity) VALUE (?,?,?)", userID, productID, quantity)
	insertedId, err := sqlResult.LastInsertId()
	return int(insertedId), err
}

func (repository CartRepositoryMySQL) UpdateCart(userID int, productID int, quantity int) error {
	sqlResult := repository.DBConnection.MustExec("UPDATE carts SET quantity=? WHERE user_id = ? AND product_id = ?", quantity, userID, productID)
	rowAffected, err := sqlResult.RowsAffected()
	if rowAffected == 0 {
		return fmt.Errorf("no any row affected , update not completed")
	}
	return err
}

func (repository CartRepositoryMySQL) DeleteCart(userID int, productID int) error {
	sqlResult := repository.DBConnection.MustExec("DELETE FROM carts WHERE user_id = ? AND product_id = ?", userID, productID)
	rowAffected, err := sqlResult.RowsAffected()
	if rowAffected == 0 {
		return fmt.Errorf("no any row affected , delete not completed")
	}
	return err
}
