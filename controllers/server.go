package controllers

import (
	"my-go-api/services"

	"github.com/gin-gonic/gin"
)

type ControllerService struct {
	router *gin.Engine
	config services.AppConfigService
	ping   IController
	user   IController
}

type IController interface {
	ResolveRouter()
}

func NewRouter() *gin.Engine {
	r := gin.Default()
	r.SetTrustedProxies([]string{"127.0.0.1"})
	return r
}

func NewControllerService(
	r *gin.Engine,
	config services.AppConfigService,
	ping PingController,
	user UserController,
) ControllerService {
	return ControllerService{
		router: r,
		config: config,
		ping:   ping,
		user:   user,
	}
}

func (c ControllerService) Start() {
	// Resolving route
	c.ping.ResolveRouter()
	c.user.ResolveRouter()

	c.router.Run() // listen and serve on 0.0.0.0:8080
}
