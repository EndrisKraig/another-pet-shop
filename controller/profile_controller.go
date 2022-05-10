package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"playground.io/another-pet-store/middleware"
	"playground.io/another-pet-store/service"
)

type ProfileController interface {
	GetProfile(c *gin.Context)
}

type SimpleProfileController struct {
	profileService service.ProfileService
}

func NewProfileController(profileService service.ProfileService) ProfileController {
	return &SimpleProfileController{profileService: profileService}
}

func (controller *SimpleProfileController) GetProfile(c *gin.Context) {
	claims, err := middleware.GetClaims(c)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Unknown server error"})
		return
	}
	userId := claims["userId"]
	profile, err := controller.profileService.GetProfile(int(userId.(float64)))
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Unknown server error"})
		return
	}
	c.IndentedJSON(http.StatusOK, profile)
}
