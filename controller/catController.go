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

	var catResponse = catService.FindAllCats(page, limit)
	c.IndentedJSON(http.StatusOK, catResponse)
}

func PostCats(c *gin.Context) {
	var newCat dto.Cat

	if err := c.BindJSON(&newCat); err != nil {
		return
	}

	newCat = catService.AddCat(&newCat)
	c.IndentedJSON(http.StatusCreated, newCat)
}

func GetCatByID(c *gin.Context) {
	id := c.Param("id")

	var cat dto.Cat = catService.FindCatById(id)
	c.IndentedJSON(http.StatusOK, cat)

}
