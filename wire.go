//go:build wireinject
// +build wireinject

package main

import (
	"my-go-api/controllers"

	"github.com/google/wire"
)

func InitializeEvent() (controllers.ControllerService, error) {
	wire.Build(
		controllers.NewRouter,
		controllers.NewControllerService,
		controllers.NewPingController,
	)
	return controllers.ControllerService{}, nil
}
