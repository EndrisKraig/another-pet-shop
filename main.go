package main

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"playground.io/another-pet-store/db"
	"playground.io/another-pet-store/model"
)

func main() {
	router := gin.Default()
	// CORS for https://foo.com and https://github.com origins, allowing:
	// - PUT and PATCH methods
	// - Origin header
	// - Credentials share
	// - Preflight requests cached for 12 hours
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"PUT", "PATCH", "GET", "POST"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "https://github.com"
		},
		MaxAge: 12 * time.Hour,
	}))
	router.GET("/cats", getCats)
	router.GET("/cats/:id", getCatByID)
	router.POST("/cats", postCats)
	router.Run("localhost:8080")
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
