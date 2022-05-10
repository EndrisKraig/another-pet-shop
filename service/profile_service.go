package service

import (
	"fmt"

	"playground.io/another-pet-store/db"
	"playground.io/another-pet-store/dto"
)

type ProfileService interface {
	CreateProfile(userId int) error
	GetProfile(userId int) (*dto.Profile, error)
	ChangeBalance(animalPrice, profileCash int) (int, error)
	ChangeBalanceFunc() func(int, int) (int, error)
}

type SimpleProfileService struct {
	profileRepository db.ProfileRepository
}

func NewProfileService(profileRepository db.ProfileRepository) ProfileService {
	return &SimpleProfileService{profileRepository: profileRepository}
}

func (service *SimpleProfileService) CreateProfile(userId int) error {
	err := service.profileRepository.CreateProfile(userId)
	if err != nil {
		return fmt.Errorf("error during create profile: %w", err)
	}
	return nil
}

func (service *SimpleProfileService) GetProfile(userId int) (*dto.Profile, error) {
	profile, err := service.profileRepository.GetProfileByUserId(int64(userId))
	if err != nil {
		return nil, err
	}
	return &dto.Profile{Id: profile.ID, Balance: profile.Balance, Image_url: profile.Image_url, Notes: profile.Notes, Nickname: profile.Nickname}, nil
}

func (service *SimpleProfileService) ChangeBalance(animalPrice, profileCash int) (int, error) {
	newBalance := profileCash - animalPrice
	if newBalance < 0 {
		return 0, fmt.Errorf("not enough money to buy an animal, you need %d more", newBalance)
	}
	return newBalance, nil
}

func (service *SimpleProfileService) ChangeBalanceFunc() func(int, int) (int, error) {
	return service.ChangeBalance
}
