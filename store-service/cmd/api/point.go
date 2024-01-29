package api

import (
	"log"
	"net/http"

	"store-service/internal/point"

	"github.com/gin-gonic/gin"
)

type PointAPI struct {
	PointService point.PointService
}

func (api PointAPI) DeductPointHandler(context *gin.Context) {
	var request point.SubmitedPoint
	if err := context.BindJSON(&request); err != nil {
		context.String(http.StatusBadRequest, err.Error())
		log.Printf("bad request %s", err.Error())
		return
	}

	uid := 1
	res, err := api.PointService.DeductPoint(uid, request)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	context.JSON(http.StatusOK, res)
}

func (api PointAPI) TotalPointHandler(context *gin.Context) {
	uid := 1
	res, err := api.PointService.TotalPoint(uid)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	context.JSON(http.StatusOK, res)
}
