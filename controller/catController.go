package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"playground.io/another-pet-store/dto"
	"playground.io/another-pet-store/service"
)

var catService service.CatServiceInstance = service.CatServiceInstance{}

func GetCats(c *gin.Context) {
	var queryParams = c.Request.URL.Query()
	var limit = "100"
	var page = "1"

	if len(queryParams["limit"]) > 0 {
		limit = queryParams["limit"][0]
	}
	if len(queryParams["page"]) > 0 {
		page = queryParams["page"][0]
	}

	catResponse, err := catService.FindAllCats(page, limit)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, catResponse)
}

func PostCats(c *gin.Context) {
	var newCat dto.Cat

	if err := c.BindJSON(&newCat); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Wrong body message"})
		return
	}

	err := catService.AddCat(&newCat)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusCreated, newCat)
}

func GetCatByID(c *gin.Context) {
	id := c.Param("id")

	cat, err := catService.FindCatById(id)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, cat)

}
