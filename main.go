package main

import (
	"log"

	"github.com/joho/godotenv"
	"playground.io/another-pet-store/controller"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	profileController := NewProfileController()
	loginController := NewLoginController()
	animalController := NewAnimalController()
	referenceController := NewReferenceController()
	controller.Init(loginController, animalController, profileController, referenceController)
}
