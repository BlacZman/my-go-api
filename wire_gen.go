// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"my-go-api/controllers"
	"my-go-api/services"
)

// Injectors from wire.go:

func InitializeEvent() (controllers.ControllerService, error) {
	engine := controllers.NewRouter()
	appConfigService := services.NewAppConfigService()
	pingController := controllers.NewPingController(engine, appConfigService)
	db := services.NewDatabaseConnection(appConfigService)
	databaseService := services.NewDatabaseService(db, appConfigService)
	userService := services.NewUserService(databaseService)
	userController := controllers.NewUserController(engine, appConfigService, userService)
	controllerService := controllers.NewControllerService(engine, appConfigService, pingController, userController)
	return controllerService, nil
}
