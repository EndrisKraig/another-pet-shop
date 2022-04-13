package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"playground.io/another-pet-store/service"
)

type SpecialOfferController interface {
	GetSpecialOffers(c *gin.Context)
}

type SimpleOfferController struct {
	specialOfferService service.SpecialOfferService
}

func NewSpecialOfferController(specialOfferService service.SpecialOfferService) SpecialOfferController {
	return &SimpleOfferController{specialOfferService: specialOfferService}
}

func (controller *SimpleOfferController) GetSpecialOffers(c *gin.Context) {
	specialOfferService := controller.specialOfferService
	offers, err := specialOfferService.GetAllActiveSpecialOffers()
	if err != nil {
		fmt.Println(err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Error during getting special offers!"})
		return
	}
	c.IndentedJSON(http.StatusOK, offers)
}
