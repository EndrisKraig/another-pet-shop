package controller

import (
	"net/http"

	"playground.io/another-pet-store/dto"
	"playground.io/another-pet-store/middleware"
	"playground.io/another-pet-store/service"

	"github.com/gin-gonic/gin"
)

var LoginControllerInstance LoginController

//login contorller interface
type LoginController interface {
	Login(ctx *gin.Context) string
}

type loginController struct {
	loginService service.LoginService
	jWtService   service.JWTService
}

func LoginHandler(loginService service.LoginService, jWtService service.JWTService) LoginController {
	return &loginController{
		loginService: loginService,
		jWtService:   jWtService,
	}
}

func (controller *loginController) Login(ctx *gin.Context) string {
	var credential dto.LoginCredentials
	err := ctx.ShouldBind(&credential)
	if err != nil {
		return "no data found"
	}
	var user = &dto.User{Username: credential.Email, Password: credential.Password}
	isUserAuthenticated := controller.loginService.LoginUser(user)
	if isUserAuthenticated {
		return controller.jWtService.GenerateToken(credential.Email, true)

	}
	return ""
}

func GetToken(ctx *gin.Context) {
	token := LoginControllerInstance.Login(ctx)
	if token != "" {
		ctx.JSON(http.StatusOK, gin.H{
			"token": token,
		})
	} else {
		ctx.JSON(http.StatusUnauthorized, nil)
	}
}

func Me(c *gin.Context) {
	i, _ := middleware.Me(c)
	c.IndentedJSON(http.StatusOK, i)
}
