package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"playground.io/another-pet-store/dto"
	"playground.io/another-pet-store/service"
)

type CatController interface {
	GetCats(c *gin.Context)
	AddCat(c *gin.Context)
	FindCatByID(c *gin.Context)
	UpdateCat(c *gin.Context)
}

type SimpleCatController struct {
	catService service.CatService
}

func (catController *SimpleCatController) GetCats(c *gin.Context) {
	var queryParams = c.Request.URL.Query()
	var limit = 100
	var page = 1

	if len(queryParams["limit"]) > 0 {
		limitParam := queryParams["limit"][0]
		var err error
		limit, err = strconv.Atoi(limitParam)
		if err != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Limit param is incorrect int"})
			return
		}
	}
	if len(queryParams["page"]) > 0 {
		pageParam := queryParams["page"][0]
		var err error
		page, err = strconv.Atoi(pageParam)
		if err != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Page param is incorrect int"})
			return
		}
	}
	var catService = catController.catService
	catResponse, err := catService.FindAllCats(page, limit)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, catResponse)
}

func (catController *SimpleCatController) AddCat(c *gin.Context) {
	var newCat dto.Cat

	if err := c.BindJSON(&newCat); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Wrong body message"})
		return
	}
	var catService = catController.catService
	err := catService.AddCat(&newCat)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusCreated, newCat)
}

func (catController *SimpleCatController) FindCatByID(c *gin.Context) {
	id := c.Param("id")
	var catService = catController.catService
	cat, err := catService.FindCatById(id)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, cat)

}

func (controller *SimpleCatController) UpdateCat(c *gin.Context) {
	id := c.Param("id")
	authHeader := c.GetHeader("Authorization")
	fmt.Println("Header: " + authHeader)
	if authHeader == "" {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	var catService = controller.catService
	err := catService.UpdateCat(id, authHeader)
	if err != nil {
		fmt.Println(err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, "OK")
}
