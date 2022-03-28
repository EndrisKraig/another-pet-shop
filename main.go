package main

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"playground.io/another-pet-store/controller"
	"playground.io/another-pet-store/db"
	"playground.io/another-pet-store/middleware"
	"playground.io/another-pet-store/model"
	"playground.io/another-pet-store/service"
)

func main() {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PUT", "PATCH", "GET", "POST", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "http://localhost:3000"
		},
		MaxAge: 12 * time.Hour,
	}))
	router.GET("/cats", getCats)
	router.GET("/cats/:id", getCatByID)
	router.POST("/cats", postCats)

	var loginService service.LoginService = service.StaticLoginService()
	var jwtService service.JWTService = service.JWTAuthService()
	var loginController controller.LoginController = controller.LoginHandler(loginService, jwtService)

	router.POST("/login", func(ctx *gin.Context) {
		token := loginController.Login(ctx)
		if token != "" {
			ctx.JSON(http.StatusOK, gin.H{
				"token": token,
			})
		} else {
			ctx.JSON(http.StatusUnauthorized, nil)
		}
	})

	router.GET("/me", middleware.AuthorizeJWT(), me)

	router.Run("localhost:8080")
}

func me(c *gin.Context) {
	i, _ := middleware.Me(c)
	c.IndentedJSON(http.StatusOK, i)
}

func getCats(c *gin.Context) {
	var cats []model.Cat = db.FindAllCats()
	c.IndentedJSON(http.StatusOK, cats)
}

func postCats(c *gin.Context) {
	var newCat model.Cat

	if err := c.BindJSON(&newCat); err != nil {
		return
	}

	db.AddCat(&newCat)
	c.IndentedJSON(http.StatusCreated, newCat)
}

func getCatByID(c *gin.Context) {
	id := c.Param("id")
	intId, err := strconv.Atoi(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "cat not found"})
		return
	}
	var cat model.Cat = db.FindCatById(intId)
	c.IndentedJSON(http.StatusOK, cat)

}
