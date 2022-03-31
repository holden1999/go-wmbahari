package api

import (
	"github.com/gin-gonic/gin"
	"go-wmb/usecase"
)

type CustomerApi struct {
	OrderUseCase   usecase.OrderUseCase
	PaymentUseCase usecase.PaymentUseCase
}

func (c *CustomerApi) NewOrder() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"Order Here :": "go-wmb",
		})
	}
}

func (c *CustomerApi) NewPayment() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"Order Here :": "go-wmb",
		})
	}
}

func NewOrderApi(foodRoute *gin.RouterGroup) {
	api := CustomerApi{}
	foodRoute.POST("/order", api.NewOrder())
}

func NewPaymentApi(foodRoute *gin.RouterGroup) {
	api := CustomerApi{}
	foodRoute.POST("/payment", api.NewOrder())
}
