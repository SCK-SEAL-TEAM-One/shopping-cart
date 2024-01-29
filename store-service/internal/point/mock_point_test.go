package point_test

import (
	"store-service/internal/point"

	"github.com/stretchr/testify/mock"
)

type mockPointRepository struct {
	mock.Mock
}

func (repo *mockPointRepository) GetPoints(userID int) ([]point.Point, error) {
	argument := repo.Called(userID)
	return argument.Get(0).([]point.Point), argument.Error(1)
}

func (repo *mockPointRepository) CreatePoint(userID int, amount int) (int, error) {
	argument := repo.Called(userID, amount)
	return argument.Int(0), argument.Error(1)
}
