package point

import (
	"fmt"
	"log"
)

type PointService struct {
	PointRepository PointRepository
}

func (pointService PointService) DeductPoint(uid int, submitedPoint SubmitedPoint) (TotalPoint, error) {
	total, err := pointService.TotalPoint(uid)
	if err != nil {
		log.Printf("pointService.TotalPoint internal error %s", err.Error())
		return TotalPoint{}, err
	}

	if submitedPoint.Amount+total.Point < 0 {
		return TotalPoint{}, fmt.Errorf("points are not enough, please try again")
	}

	_, err_ := pointService.PointRepository.CreatePoint(uid, submitedPoint.Amount)
	if err_ != nil {
		log.Printf("PointRepository.CreatePoint internal error %s", err.Error())
		return TotalPoint{}, err_
	}
	return pointService.TotalPoint(uid)
}

func (pointService PointService) TotalPoint(uid int) (TotalPoint, error) {
	points, err := pointService.PointRepository.GetPoints(uid)
	if err != nil {
		log.Printf("PointRepository.GetPoints internal error %s", err.Error())
	}

	total := 0
	for _, point := range points {
		total += point.Amount
	}
	return TotalPoint{
		Point: total,
	}, err
}
