package service

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
	"playground.io/another-pet-store/dto"
)

type LoginService interface {
	LoginUser(user *dto.User) (string, error)
	NewUser(user *dto.User)
}

type SimpleLoginService struct {
	userService    UserService
	profileService ProfileService
	jwtService     JWTService
}

func NewLoginService(userService UserService, profileService ProfileService, jwtService JWTService) LoginService {
	return &SimpleLoginService{userService: userService, profileService: profileService, jwtService: jwtService}
}

func (loginService *SimpleLoginService) LoginUser(user *dto.User) (string, error) {
	var dbUser = loginService.userService.FindUserByUsername(user.Username)
	profile, err := loginService.profileService.GetProfile(int(dbUser.ID))
	if err != nil {
		return "", fmt.Errorf("profile wasn't found: %w", err)
	}
	var isPasswordsEquals = checkPasswordHash(user.Password, dbUser.Hash)
	if !isPasswordsEquals {
		return "", fmt.Errorf("password not equals")
	}
	token := loginService.jwtService.GenerateToken(profile.Nickname, true, int(dbUser.ID), int(profile.Id))
	return token, nil
}

func (loginService *SimpleLoginService) NewUser(user *dto.User) {
	loginService.userService.RegisterUser(user)
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	fmt.Println(err)
	return err == nil
}
