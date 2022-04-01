package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"playground.io/another-pet-store/dto"
	"playground.io/another-pet-store/service"
)

func PostUser(c *gin.Context) {
	var newUser dto.User

	if err := c.BindJSON(&newUser); err != nil {
		return
	}

	service.RegisterUser(&newUser)
	c.IndentedJSON(http.StatusCreated, newUser)
}
