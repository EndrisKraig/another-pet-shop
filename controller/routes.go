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

	router.POST("/login", GetToken)

	router.GET("/cats", GetCats)
	router.GET("/cats/:id", GetCatByID)
	router.POST("/cats", PostCats)

	router.POST("/user", PostUser)

	var loginService service.LoginService = service.DbLoginService()
	var jwtService service.JWTService = service.JWTAuthService()
	//TODO how to properly deal with services?
	LoginControllerInstance = LoginHandler(loginService, jwtService)

	router.GET("/me", middleware.AuthorizeJWT(), Me)

	router.Run("localhost:8080")
}
