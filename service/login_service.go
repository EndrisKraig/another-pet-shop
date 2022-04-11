package service

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
	"playground.io/another-pet-store/dto"
)

type LoginService interface {
	LoginUser(user *dto.User) bool
	NewUser(user *dto.User)
}

type SimpleLoginService struct {
	userService UserService
}

func NewLoginService(userService UserService) LoginService {
	return &SimpleLoginService{userService: userService}
}

func (loginService *SimpleLoginService) LoginUser(user *dto.User) bool {
	var dbUser = loginService.userService.FindUserByUsername(user.Username)
	var isIt = checkPasswordHash(user.Password, dbUser.Hash)
	return isIt
}

func (loginService *SimpleLoginService) NewUser(user *dto.User) {
	loginService.userService.RegisterUser(user)
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	fmt.Println(err)
	return err == nil
}
