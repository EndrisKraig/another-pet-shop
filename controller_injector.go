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
	wire.Build(controller.NewAnimalController, service.NewAnimalService, db.NewAnimalRepository, service.NewProfileService, db.NewProfileRepository)
	return &controller.SimpleAnimalController{}
}

func NewLoginController() controller.LoginController {
	wire.Build(controller.NewLoginController, service.NewLoginService, service.NewJWTService, service.NewUserService, service.NewProfileService, db.NewProfileRepository)
	return &controller.SimpleLoginController{}
}

func NewProfileController() controller.ProfileController {
	wire.Build(controller.NewProfileController, service.NewProfileService, db.NewProfileRepository)
	return &controller.SimpleProfileController{}
}

func NewReferenceController() controller.ReferenceController {
	wire.Build(controller.NewReferenceController, service.NewReferenceService, db.NewReferenceRepository)
	return &controller.SimpleReferenceController{}
}

func NewSpecialOfferController() controller.SpecialOfferController {
	wire.Build(controller.NewSpecialOfferController, service.NewSpecialOfferServie, db.NewSpecialOfferRepository)
	return &controller.SimpleOfferController{}
}
