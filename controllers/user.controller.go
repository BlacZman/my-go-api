package controllers

import (
	"errors"
	"my-go-api/services"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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

func (c UserController) getUser(context *gin.Context) {
	// Setup input from context
	idString := context.Param("id")
	id, err := strconv.ParseUint(idString, 10, 0)

	// Binding & Validating
	if err != nil {
		errorMessage := err.Error()
		context.JSON(400, gin.H{"error": errorMessage})
		return
	}

	// Business Logic
	user, err := c.userService.GetUser(uint(id))

	// Handle error from business logic
	if err != nil {
		statusCode := 500
		errorMessage := err.Error()

		if errors.Is(err, gorm.ErrRecordNotFound) {
			statusCode = 404
		}

		context.JSON(statusCode, gin.H{"error": errorMessage})
		return
	}

	// Compile successful response
	result := gin.H{
		"data": user,
	}

	// Send result
	context.JSON(200, result)
}

func (c UserController) createUser(context *gin.Context) {
	// Binding & Validating
	var newUser services.CreateUserBody
	if err := context.ShouldBindJSON(&newUser); err != nil {
		context.JSON(400, gin.H{"error": err})
		return
	}

	// Business Logic
	user, err := c.userService.CreateUser(newUser)

	// Handle error from business logic
	if err != nil {
		statusCode := 500
		errorMessage := err.Error()

		context.JSON(statusCode, gin.H{"error": errorMessage})
		return
	}

	// Compile successful response
	result := gin.H{
		"data": gin.H{
			"id": user.ID,
		},
	}

	// Send result
	context.JSON(200, result)
}

func (c UserController) ResolveRouter() {
	c.router.GET("/user/:id", c.getUser)
	c.router.POST("/user", c.createUser)
}
