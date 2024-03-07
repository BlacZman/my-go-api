package controllers

import (
	"my-go-api/services"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	router      *gin.Engine
	config      services.AppConfigService
	userService services.UserService
}

func NewUserController(router *gin.Engine, config services.AppConfigService, userService services.UserService) UserController {
	return UserController{
		router:      router,
		config:      config,
		userService: userService,
	}
}

type UriUser struct {
	id uint `uri:"id" binding:"required"`
}

func (c UserController) getUser(context *gin.Context) {
	var findUser UriUser
	if err := context.ShouldBindUri(&findUser); err != nil {
		context.JSON(400, gin.H{"msg": err})
		return
	}
	user := c.userService.GetUser(findUser.id)
	result := gin.H{
		"data": user,
	}
	context.JSON(200, result)
}

func (c UserController) ResolveRouter() {
	c.router.GET("/user/:id", c.getUser)
}
