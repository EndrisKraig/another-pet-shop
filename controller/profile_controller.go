package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
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
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	var profile, err = controller.profileService.GetProfile(authHeader)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Unknown server error"})
		return
	}
	c.IndentedJSON(http.StatusOK, profile)
}
