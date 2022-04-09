package controller

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"playground.io/another-pet-store/db"
	"playground.io/another-pet-store/middleware"
	"playground.io/another-pet-store/service"
)

func Init() {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PUT", "PATCH", "GET", "POST", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "http://localhost:3000"
		},
		MaxAge: 12 * time.Hour,
	}))

	var userService service.UserService = service.CreateUserService()
	var loginService service.LoginService = service.DbLoginService(userService)
	var jwtService service.JWTService = service.JWTAuthService()

	loginController := LoginHandler(loginService, jwtService)
	profileRepository := db.CreateProfileRepository()
	profileService := service.CreateProfileService(userService, jwtService, profileRepository)
	ProfileController := createProfileController(profileService)
	animalRepository := db.CreateAnimalRepository()
	var animalService service.AnimalService = service.CreateAnimalService(profileService, animalRepository)
	var animalController = &SimpleAnimalController{animalService: animalService}
	router.POST("/login", loginController.Login)
	router.POST("/user", loginController.AddUser)
	router.GET("/me", middleware.AuthorizeJWT(), loginController.Me)
	router.GET("/profile", middleware.AuthorizeJWT(), ProfileController.GetProfile)
	router.GET("/animals", animalController.GetAnimals)
	router.GET("/animals/:id", animalController.FindAnimalByID)
	router.POST("/animals", animalController.AddAnimal)
	router.POST("/animals/:id", animalController.UpdateAnimal)

	router.Run("localhost:8080")
}
