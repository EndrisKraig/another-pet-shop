package controller

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"playground.io/another-pet-store/chat"
	"playground.io/another-pet-store/middleware"
)

func Init(loginController LoginController, animalController AnimalController, profileController ProfileController, referenceController ReferenceController, specialOfferController SpecialOfferController) {
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

	hub := chat.NewHub()
	go hub.Run()

	router.POST("/login", loginController.Login)
	router.POST("/user", loginController.AddUser)
	router.GET("/me", middleware.AuthorizeJWT(), loginController.Me)
	router.GET("/profile", middleware.AuthorizeJWT(), profileController.GetProfile)
	router.GET("/animals", animalController.GetAnimals)
	router.GET("/animals/specials", specialOfferController.GetSpecialOffers)
	router.GET("/animals/:id", animalController.FindAnimalByID)
	router.POST("/animals", animalController.AddAnimal)
	router.POST("/animals/:id", middleware.AuthorizeJWT(), animalController.UpdateAnimal)
	router.GET("/references", referenceController.GetReferences)
	router.GET("/chat", func(c *gin.Context) { chat.ServeWs(hub, c.Writer, c.Request) })
	router.Run("localhost:8080")
}
