package point

import (
	"log"
)

type PointService struct {
	PointRepository PointRepository
}

func (pointService PointService) DeductPoint(uid int, submitedPoint SubmitedPoint) (TotalPoint, error) {
	_, err := pointService.PointRepository.CreatePoint(uid, submitedPoint.Amount)
	if err != nil {
		log.Printf("PointRepository.CreatePoint internal error %s", err.Error())
		return TotalPoint{}, err
	}
	res, err_ := pointService.TotalPoint(uid)
	return res, err_
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
