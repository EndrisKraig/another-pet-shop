package main

import (
	"github.com/joho/godotenv"
	"playground.io/another-pet-store/controller"
	"playground.io/another-pet-store/logs"
)

func init() {
	if err := godotenv.Load(); err != nil {
		logs.Logger.Error("No .env file found")
	}
}

func main() {
	logs.InitLogger()
	logs.Logger.Info("Application started")
	profileController := NewProfileController()
	loginController := NewLoginController()
	animalController := NewAnimalController()
	referenceController := NewReferenceController()
	specialOfferController := NewSpecialOfferController()
	chatController := NewChatController()
	controller.Init(loginController, animalController, profileController, referenceController, specialOfferController, chatController)
}
