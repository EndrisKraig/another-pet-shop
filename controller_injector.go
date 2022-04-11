//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"playground.io/another-pet-store/controller"
	"playground.io/another-pet-store/db"
	"playground.io/another-pet-store/service"
)

//TODO cross reference is a sign of bad design? Or there is another way to make exactly one instance of a service

func NewAnimalController() controller.AnimalController {
	wire.Build(controller.NewAnimalController, service.NewAnimalService, db.NewAnimalRepository, service.NewProfileService, service.NewJWTService, service.NewUserService, db.NewProfileRepository)
	return &controller.SimpleAnimalController{}
}

func NewLoginController() controller.LoginController {
	wire.Build(controller.NewLoginController, service.NewLoginService, service.NewJWTService, service.NewUserService)
	return &controller.SimpleLoginController{}
}

func NewProfileController() controller.ProfileController {
	wire.Build(controller.NewProfileController, service.NewProfileService, service.NewUserService, service.NewJWTService, db.NewProfileRepository)
	return &controller.SimpleProfileController{}
}
