package controller

import (
	"fmt"
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

type catsResponse struct {
	Cats    []model.Cat `json:"cats"`
	MaxPage int         `json:"maxPage"`
}

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
	fmt.Println("limit " + limit + " page " + page)
	cats, allCatsCount := db.FindAllCats(page, limit)
	limitInt, _ := strconv.ParseInt(limit, 10, 64)
	//TODO proper calculation in service, e.g. there is no additional page when all%limit = 0
	maxPage := allCatsCount/limitInt + 1
	c.IndentedJSON(http.StatusOK, catsResponse{Cats: cats, MaxPage: int(maxPage)})
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
