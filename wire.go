//go:build wireinject
// +build wireinject

package main

import (
	"my-go-api/controllers"
	"my-go-api/services"

	"github.com/google/wire"
)

func InitializeEvent() (controllers.ControllerService, error) {
	wire.Build(
		services.NewAppConfigService,
		services.NewDatabaseConnection,
		services.NewDatabaseService,
		services.NewUserService,
		controllers.NewRouter,
		controllers.NewUserController,
		controllers.NewControllerService,
		controllers.NewPingController,
	)
	return controllers.ControllerService{}, nil
}
