package point_test

import (
	"store-service/internal/point"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_DeductPoint_Input_Amount_100_Should_be_Point_100(t *testing.T) {
	expected := point.TotalPoint{
		Point: 100,
	}
	uid := 1
	res := []point.Point{
		{
			ID:     1,
			UserID: 1,
			Amount: 100,
		},
	}

	mockPointRepository := new(mockPointRepository)
	mockPointRepository.On("CreatePoint", uid, 100).Return(1, nil)
	mockPointRepository.On("GetPoints", uid).Return(res, nil)

	pointService := point.PointService{
		PointRepository: mockPointRepository,
	}
	actual, err := pointService.DeductPoint(uid, point.SubmitedPoint{
		Amount: 100,
	})

	assert.Equal(t, expected, actual)
	assert.Equal(t, nil, err)
}

func Test_TotalPoint_Point_100_and_50_Should_be_Point_150(t *testing.T) {
	expected := point.TotalPoint{
		Point: 150,
	}
	uid := 1
	res := []point.Point{
		{
			ID:     1,
			UserID: 1,
			Amount: 100,
		},
		{
			ID:     2,
			UserID: 1,
			Amount: 50,
		},
	}

	mockPointRepository := new(mockPointRepository)
	mockPointRepository.On("GetPoints", uid).Return(res, nil)

	pointService := point.PointService{
		PointRepository: mockPointRepository,
	}
	actual, err := pointService.TotalPoint(uid)

	assert.Equal(t, expected, actual)
	assert.Equal(t, nil, err)
}

func Test_TotalPoint_Point_100_and_Minus_50_Should_be_Point_50(t *testing.T) {
	expected := point.TotalPoint{
		Point: 50,
	}
	uid := 1
	res := []point.Point{
		{
			ID:     1,
			UserID: 1,
			Amount: 100,
		},
		{
			ID:     2,
			UserID: 1,
			Amount: -50,
		},
	}

	mockPointRepository := new(mockPointRepository)
	mockPointRepository.On("GetPoints", uid).Return(res, nil)

	pointService := point.PointService{
		PointRepository: mockPointRepository,
	}
	actual, err := pointService.TotalPoint(uid)

	assert.Equal(t, expected, actual)
	assert.Equal(t, nil, err)
}
