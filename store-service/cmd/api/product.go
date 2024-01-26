package api

import (
	"net/http"
	"store-service/internal/product"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductAPI struct {
	ProductRepository product.ProductRepository
}

func (api ProductAPI) SearchHandler(context *gin.Context) {
	keyword := context.DefaultQuery("q", "")
	limit := context.DefaultQuery("limit", "30")
	offset := context.DefaultQuery("offset", "0")

	productResult, err := api.ProductRepository.GetProducts(keyword, limit, offset)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	context.JSON(http.StatusOK, productResult)
}

func (api ProductAPI) GetProductHandler(context *gin.Context) {
	idParam := context.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": "id is not integer",
		})
		return
	}
	product, err := api.ProductRepository.GetProductByID(id)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	context.JSON(http.StatusOK, product)
}
