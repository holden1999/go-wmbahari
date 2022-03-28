package main

import (
	"github.com/gin-gonic/gin"
	"go-wmb/config"
	"go-wmb/delivery/api"
)

type AppServer interface {
	Run()
}

type appServer struct {
	routerEngine *gin.Engine
	config       *config.Config
}

func (s *appServer) initHandlers() {
	s.v1()
}

func (s *appServer) v1() {
	CustomerApiGroup := s.routerEngine.Group("/customer")
	api.NewOrderApi(CustomerApiGroup)

	FoodListApiGroup := s.routerEngine.Group("/list")
	api.GetListApi(FoodListApiGroup)
}

func (s *appServer) Run() {
	s.initHandlers()
	err := s.routerEngine.Run("localhost:3333")
	if err != nil {
		panic(err)
	}
}

func Server() AppServer {
	r := gin.Default()
	c := config.NewConfig()
	return &appServer{
		routerEngine: r,
		config:       c,
	}
}
