package api

import (
	"github.com/gin-gonic/gin"
	"go-wmb/usecase"
)

type FoodApi struct {
	CustomerOrderUseCase usecase.CustomerOrderUseCase
}

func (o *FoodApi) NewOrder() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"Order Here :": "go-wmb",
		})
	}
}

func (o *FoodApi) GetList() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"Food List :": "go-wmb",
		})
	}
}

func NewOrderApi(foodRoute *gin.RouterGroup) {
	api := FoodApi{}
	foodRoute.POST("/order", api.NewOrder())
}

func GetListApi(foodRoute *gin.RouterGroup) {
	api := FoodApi{}
	foodRoute.GET("/get", api.GetList())
}
