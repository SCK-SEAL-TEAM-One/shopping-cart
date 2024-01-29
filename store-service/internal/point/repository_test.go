//go:build integration
// +build integration

package point_test

import (
	"store-service/internal/point"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
)

func Test_PointRepository(t *testing.T) {
	connection, err := sqlx.Connect("mysql", "sealteam:sckshuhari@(localhost:3306)/toy")
	if err != nil {
		t.Fatalf("cannot tearup data err %s", err)
	}
	repository := point.PointRepositoryMySQL{
		DBConnection: connection,
	}

	t.Run("CreatePoint_Input_Amount_100_Should_Be_PointID_No_Error", func(t *testing.T) {
		uid, amount := 1, 100
		actualId, err := repository.CreatePoint(uid, amount)

		assert.Equal(t, nil, err)
		assert.NotEmpty(t, actualId)
	})

	t.Run("CreatePoint_Input_Amount_Minus_30_Should_Be_PointID_No_Error", func(t *testing.T) {
		uid, amount := 1, -30
		actualId, err := repository.CreatePoint(uid, amount)

		assert.Equal(t, nil, err)
		assert.NotEmpty(t, actualId)
	})

	t.Run("GetPoints_Should_Be_Length_No_Error", func(t *testing.T) {
		uid := 1

		actualPoints, err := repository.GetPoints(uid)
		assert.Len(t, actualPoints, 2)
		assert.Equal(t, nil, err)
	})
}
