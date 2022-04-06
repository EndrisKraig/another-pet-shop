package controller

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
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

	var catService service.CatService = service.CatServiceInstance{}
	var catController = &SimpleCatController{catService: catService}

	router.GET("/cats", catController.GetCats)
	router.GET("/cats/:id", catController.FindCatByID)
	router.POST("/cats", catController.AddCat)

	var userService service.UserService = service.CreateUserService()
	var loginService service.LoginService = service.DbLoginService(userService)
	var jwtService service.JWTService = service.JWTAuthService()

	loginController := LoginHandler(loginService, jwtService)

	router.POST("/login", loginController.Login)
	router.POST("/user", loginController.AddUser)
	router.GET("/me", middleware.AuthorizeJWT(), loginController.Me)

	router.Run("localhost:8080")
}
