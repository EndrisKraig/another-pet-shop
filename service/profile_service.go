package service

import (
	"fmt"

	"github.com/dgrijalva/jwt-go"
	"playground.io/another-pet-store/db"
	"playground.io/another-pet-store/dto"
	"playground.io/another-pet-store/model"
)

type ProfileService interface {
	CreateProfile(token string) error
	GetProfile(token string) (*dto.Profile, error)
	ChangeBalance(value float32, profile *dto.Profile) error
}

type SimpleProfileService struct {
	userService       UserService
	jwtService        JWTService
	profileRepository db.ProfileRepository
}

func NewProfileService(userService UserService, jwtServices JWTService, profileRepository db.ProfileRepository) ProfileService {
	return &SimpleProfileService{userService: userService, jwtService: jwtServices, profileRepository: profileRepository}
}

func (service *SimpleProfileService) CreateProfile(token string) error {
	var user, err = service.getUserFromToken(token)
	if err != nil {
		return fmt.Errorf("Error during get user: %w", err)
	}
	err = service.profileRepository.CreateProfile(nil, int(user.ID))
	if err != nil {
		return fmt.Errorf("Error during create profile: %w", err)
	}
	return nil
}

func (service *SimpleProfileService) GetProfile(token string) (*dto.Profile, error) {
	var user, err = service.getUserFromToken(token)
	if err != nil {
		return nil, fmt.Errorf("Error during get user: %w", err)
	}

	fmt.Println(user.ID)
	profile, err := service.profileRepository.GetProfileByUserId(user.ID)
	if err != nil {
		service.CreateProfile(token)
		profile, err = service.profileRepository.GetProfileByUserId(user.ID)
		fmt.Println(err)
	}
	return &dto.Profile{Id: profile.ID, Balance: profile.Balance, Image_url: profile.Image_url, Notes: profile.Notes, Nickname: profile.Nickname}, nil
}

func (service *SimpleProfileService) getUserFromToken(token string) (*model.User, error) {
	var claims, err = service.jwtService.ValidateToken(token)
	if err != nil {
		return nil, fmt.Errorf("Couldn't validate token: %w", err)
	}
	var newClaims = claims.Claims.(jwt.MapClaims)
	var username string = newClaims["name"].(string)
	return service.userService.FindUserByUsername(username), nil
}

func (service *SimpleProfileService) ChangeBalance(value float32, profile *dto.Profile) error {
	newBalance := profile.Balance - float32(value)
	err := service.profileRepository.UpdateBalance(profile.Id, float64(newBalance))
	if err != nil {
		return fmt.Errorf("Error changing balance: %w", err)
	}
	return nil
}
