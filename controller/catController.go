package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"playground.io/another-pet-store/db"
	"playground.io/another-pet-store/middleware"
	"playground.io/another-pet-store/model"
)

func Me(c *gin.Context) {
	i, _ := middleware.Me(c)
	c.IndentedJSON(http.StatusOK, i)
}

func GetCats(c *gin.Context) {
	var cats []model.Cat = db.FindAllCats()
	c.IndentedJSON(http.StatusOK, cats)
}

func PostCats(c *gin.Context) {
	var newCat model.Cat

	if err := c.BindJSON(&newCat); err != nil {
		return
	}

	db.AddCat(&newCat)
	c.IndentedJSON(http.StatusCreated, newCat)
}

func GetCatByID(c *gin.Context) {
	id := c.Param("id")
	intId, err := strconv.Atoi(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "cat not found"})
		return
	}
	var cat model.Cat = db.FindCatById(intId)
	c.IndentedJSON(http.StatusOK, cat)

}
