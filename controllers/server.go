package controllers

import "github.com/gin-gonic/gin"

type ControllerService struct {
	router *gin.Engine
	ping   PingController
}

func NewRouter() *gin.Engine {
	return gin.Default()
}

func NewControllerService(r *gin.Engine, ping PingController) ControllerService {
	return ControllerService{
		router: r,
		ping:   ping,
	}
}

func (c ControllerService) Start() {
	// Resolving route
	c.ping.ResolveRouter()

	c.router.Run() // listen and serve on 0.0.0.0:8080
}
