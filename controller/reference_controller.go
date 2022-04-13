package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"playground.io/another-pet-store/service"
)

type ReferenceController interface {
	GetReferences(c *gin.Context)
}

type SimpleReferenceController struct {
	referenceService service.ReferenceService
}

func NewReferenceController(referenceService service.ReferenceService) ReferenceController {
	return &SimpleReferenceController{referenceService: referenceService}
}

func (controller *SimpleReferenceController) GetReferences(c *gin.Context) {
	queryParams := c.Request.URL.Query()
	var name string
	if len(queryParams["name"]) > 0 {
		name = queryParams["name"][0]
	} else {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Param value is required"})
		return
	}

	references, err := controller.referenceService.GetReferences(name)

	if err != nil {
		fmt.Println(err)
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "No references for " + name})
		return
	}
	c.IndentedJSON(http.StatusOK, references)
}
