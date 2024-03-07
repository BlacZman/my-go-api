package controllers

import (
	"my-go-api/services"

	"github.com/gin-gonic/gin"
)

type PingController struct {
	router *gin.Engine
	config services.AppConfigService
}

func NewPingController(router *gin.Engine, config services.AppConfigService) PingController {
	return PingController{
		router: router,
		config: config,
	}
}

func (c PingController) pingHandler(context *gin.Context) {
	result := gin.H{
		"message": "pong",
	}
	context.JSON(200, result)
}

func (c PingController) ResolveRouter() {
	c.router.GET("/ping", c.pingHandler)
}
