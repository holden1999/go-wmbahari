package api

import (
	"github.com/gin-gonic/gin"
	"go-wmb/usecase"
)

type ListApi struct {
	ListFoodUseCase  usecase.FoodUseCase
	ListTableUseCase usecase.TableUseCase
}

func (l *ListApi) GetList() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"Food List :": "go-wmb",
		})
	}
}

func GetListApi(foodRoute *gin.RouterGroup) {
	api := ListApi{}
	foodRoute.GET("/get", api.GetList())
}
