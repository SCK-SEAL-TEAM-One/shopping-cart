package payment_test

import (
	"store-service/internal/payment"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
)

func Test_PaymentRepository(t *testing.T) {
	connection, err := sqlx.Connect("mysql", "sealteam:sckshuhari@(localhost:3306)/toy")
	if err != nil {
		t.Fatalf("cannot tearup data err %s", err)
	}
	repository := payment.PaymentRepositoryMySQL{
		DBConnection: connection,
	}

	t.Run("UpdateStock_Input_Product_ID_2_No_Error", func(t *testing.T) {
		productID := 2
		quantity := 1
		err := repository.UpdateStock(productID, quantity)

		assert.Equal(t, nil, err)
	})
}
