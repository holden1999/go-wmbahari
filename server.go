package main

import (
	"github.com/gin-gonic/gin"
	"go-wmb/config"
	"go-wmb/delivery/api"
)

type AppServer interface {
	Run()
}

type customerServer struct {
	routerEngine *gin.Engine
	config       *config.Config
}

type listServer struct {
	routerEngine *gin.Engine
	config       *config.Config
}

func (c *customerServer) initHandlers() {
	c.v1()
}

func (c *customerServer) v1() {
	CustomerApiGroup := c.routerEngine.Group("/customer")
	api.NewOrderApi(CustomerApiGroup)

	ListApiGroup := c.routerEngine.Group("/list")
	api.GetListApi(ListApiGroup)

}

func (c *customerServer) Run() {
	c.initHandlers()
	api := c.config.Get("wmbahari.api.url")
	err := c.routerEngine.Run(api)
	if err != nil {
		panic(err)
	}
}

func Server() AppServer {
	r := gin.Default()
	c := config.NewConfig(".", "config.yaml")
	return &customerServer{
		routerEngine: r,
		config:       c,
	}
}
