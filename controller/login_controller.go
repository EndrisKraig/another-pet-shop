package controller

import (
	"fmt"
	"net/http"

	"playground.io/another-pet-store/dto"
	"playground.io/another-pet-store/middleware"
	"playground.io/another-pet-store/service"

	"github.com/gin-gonic/gin"
)

type LoginController interface {
	Login(ctx *gin.Context)
	AddUser(ctx *gin.Context)
	Me(ctx *gin.Context)
}

type SimpleLoginController struct {
	loginService service.LoginService
}

func NewLoginController(loginService service.LoginService) LoginController {
	return &SimpleLoginController{
		loginService: loginService,
	}
}

func (controller *SimpleLoginController) Login(c *gin.Context) {
	var credential dto.LoginCredentials
	err := c.ShouldBind(&credential)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Wrong request body"})
		return
	}
	var user = &dto.User{Username: credential.Email, Password: credential.Password}
	token, err := controller.loginService.LoginUser(user)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Authentication failed"})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"token": token})

}

func (controller *SimpleLoginController) Me(c *gin.Context) {
	payload, ok := c.Get(middleware.AuthorizationPayloadKey)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "User unauthorized"})
	}
	c.IndentedJSON(http.StatusOK, payload)
}

func (controller *SimpleLoginController) AddUser(c *gin.Context) {
	var newUser dto.User
	if err := c.BindJSON(&newUser); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Wrong request body"})
		return
	}

	controller.loginService.NewUser(&newUser)
	c.IndentedJSON(http.StatusCreated, newUser)
}
