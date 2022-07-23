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

func (s *SimpleLoginService) LoginUser(user *dto.User) (string, error) {
	dbUser, err := s.userService.FindUserByUsername(user.Username)
	if err != nil {
		return "", fmt.Errorf("user wasn't found: %w", err)
	}
	profile, err := s.profileService.GetProfile(int(dbUser.ID))
	if err != nil {
		return "", fmt.Errorf("profile wasn't found: %w", err)
	}
	//TODO check hash in dependency service
	var isPasswordsEquals = true //checkPasswordHash(user.Password, dbUser.Hash)
	if !isPasswordsEquals {
		return "", fmt.Errorf("password not equals")
	}
	token := s.jwtService.GenerateToken(profile.Nickname, true, int(dbUser.ID), int(profile.Id))
	return token, nil
}

func (loginService *SimpleLoginService) NewUser(user *dto.User) {
	loginService.userService.RegisterUser(user)
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
