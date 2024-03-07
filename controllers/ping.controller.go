package controllers

import "github.com/gin-gonic/gin"

type PingController struct {
	router *gin.Engine
}

func NewPingController(router *gin.Engine) PingController {
	return PingController{router: router}
}

func pingHandler(context *gin.Context) {
	result := gin.H{
		"message": "pong",
	}
	context.JSON(200, result)
}

func (c PingController) ResolveRouter() {
	c.router.GET("/ping", pingHandler)
}
