//+build integration

package product_test

import (
	"store-service/internal/product"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
)

func Test_ProductRepository(t *testing.T) {
	connection, err := sqlx.Connect("mysql", "sealteam:sckshuhari@(localhost:3306)/toy")
	if err != nil {
		t.Fatalf("cannot tearup data err %s", err)
	}
	repository := product.ProductRepositoryMySQL{
		DBConnection: connection,
	}

	t.Run("GetProductByID_Input_ID_2_Should_Be_Product_Detail_No_Error", func(t *testing.T) {
		expected := product.ProductDetail{
			ID:       2,
			Name:     "43 Piece dinner Set",
			Price:    12.95,
			Quantity: 10,
			Brand:    "CoolKidz",
			Image:    "/43_Piece_dinner_Set.png",
		}
		ID := 2

		actualProduct, err := repository.GetProductByID(ID)
		assert.Equal(t, expected, actualProduct)
		assert.Equal(t, err, nil)
	})

	t.Run("UpdateStock_Input_Product_ID_2_No_Error", func(t *testing.T) {
		productID := 2
		quantity := 1
		err := repository.UpdateStock(productID, quantity)

		assert.Equal(t, nil, err)
	})
}
