//+build integration

package order_test

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
	"store-service/internal/order"
	"testing"
)

func Test_OrderRepository(t *testing.T) {
	connection, err := sqlx.Connect("mysql", "sealteam:sckshuhari@(localhost:3306)/toy")
	if err != nil {
		t.Fatal("cannot tearup data")
	}
	repository := order.OrderRepositoryMySQL{
		DBConnection: connection,
	}

	t.Run("CreateOrder_Input_TotalPrice_14_dot_95_Should_Be_OrderID_1337620837", func(t *testing.T) {
		expectedId := 1337620837
		totalPrice := 14.95

		actualId, err := repository.CreateOrder(totalPrice)

		assert.Equal(t, nil, err)
		assert.Equal(t, expectedId, actualId)
	})
}
