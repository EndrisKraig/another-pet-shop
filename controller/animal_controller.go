package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"playground.io/another-pet-store/dto"
	"playground.io/another-pet-store/service"
)

type AnimalController interface {
	GetAnimals(c *gin.Context)
	AddAnimal(c *gin.Context)
	FindAnimalByID(c *gin.Context)
	UpdateAnimal(c *gin.Context)
}

type SimpleAnimalController struct {
	animalService service.AnimalService
}

func (animalController *SimpleAnimalController) GetAnimals(c *gin.Context) {
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
	var animalService = animalController.animalService
	animalResponse, err := animalService.FindAllAnimals(page, limit)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, animalResponse)
}

func (animalController *SimpleAnimalController) AddAnimal(c *gin.Context) {
	var newAnimal dto.Animal

	if err := c.BindJSON(&newAnimal); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Wrong body message"})
		return
	}
	var animalService = animalController.animalService
	err := animalService.AddAnimal(&newAnimal)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusCreated, newAnimal)
}

func (animalController *SimpleAnimalController) FindAnimalByID(c *gin.Context) {
	id := c.Param("id")
	var animalService = animalController.animalService
	animal, err := animalService.FindAnimalById(id)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, animal)

}

func (controller *SimpleAnimalController) UpdateAnimal(c *gin.Context) {
	id := c.Param("id")
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	var animalService = controller.animalService
	err := animalService.UpdateAnimal(id, authHeader)
	if err != nil {
		fmt.Println(err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, "OK")
}
