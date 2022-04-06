package controller

import (
	"net/http"

	"playground.io/another-pet-store/dto"
	"playground.io/another-pet-store/middleware"
	"playground.io/another-pet-store/service"

	"github.com/gin-gonic/gin"
)

var LoginControllerInstance LoginController

type LoginController interface {
	Login(ctx *gin.Context)
	AddUser(ctx *gin.Context)
	Me(ctx *gin.Context)
}

type loginController struct {
	loginService service.LoginService
	jWtService   service.JWTService
	userService  service.UserService
}

func LoginHandler(loginService service.LoginService, jWtService service.JWTService) *loginController {
	return &loginController{
		loginService: loginService,
		jWtService:   jWtService,
	}
}

func (controller *loginController) Login(c *gin.Context) {
	var credential dto.LoginCredentials
	err := c.ShouldBind(&credential)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Wrong request body"})
		return
	}
	var user = &dto.User{Username: credential.Email, Password: credential.Password}
	isUserAuthenticated := controller.loginService.LoginUser(user)
	if isUserAuthenticated {
		var token = controller.jWtService.GenerateToken(user.Username, true)
		c.IndentedJSON(http.StatusOK, gin.H{"token": token})
		return
	}
	c.JSON(http.StatusUnauthorized, gin.H{"message": "Authentication failed"})
}

func (controller *loginController) Me(c *gin.Context) {
	i, _ := middleware.Me(c)
	c.IndentedJSON(http.StatusOK, i)
}

func (controller *loginController) AddUser(c *gin.Context) {
	var newUser dto.User
	if err := c.BindJSON(&newUser); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Wrong request body"})
		return
	}

	controller.loginService.NewUser(&newUser)
	c.IndentedJSON(http.StatusCreated, newUser)
}
