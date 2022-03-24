package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type cat struct {
	ID       string  `json:"id"`
	Nickname string  `json:"nickname"`
	Breed    string  `json:"breed"`
	Price    float64 `json:"price"`
}

var cats = []cat{
	{ID: "1", Nickname: "Fluffy", Breed: "Persian", Price: 100},
	{ID: "2", Nickname: "Claus", Breed: "Sphynx", Price: 100},
	{ID: "3", Nickname: "Desmond", Breed: "Siberian", Price: 100},
}

func main() {
	router := gin.Default()
	router.GET("/cats", getCats)
	router.GET("/cats/:id", getCatByID)
	router.POST("/cats", postCats)
	router.Run("localhost:8080")
}

func getCats(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, cats)
}

func postCats(c *gin.Context) {
	var newCat cat

	if err := c.BindJSON(&newCat); err != nil {
		return
	}

	cats = append(cats, newCat)
	c.IndentedJSON(http.StatusCreated, newCat)
}

func getCatByID(c *gin.Context) {
	id := c.Param("id")

	for _, cat := range cats {
		if cat.ID == id {
			c.IndentedJSON(http.StatusOK, cat)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "cat not found"})
}
