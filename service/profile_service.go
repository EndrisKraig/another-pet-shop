package service

import (
	"fmt"

	"playground.io/another-pet-store/db"
	"playground.io/another-pet-store/dto"
)

type ProfileService interface {
	CreateProfile(userId int) error
	GetProfile(userId int) (*dto.Profile, error)
	ChangeBalance(value float32, profile *dto.Profile) error
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
	//TODO move create profile logic to addUser method in user service
	if err != nil {
		return nil, err
	}
	return &dto.Profile{Id: profile.ID, Balance: profile.Balance, Image_url: profile.Image_url, Notes: profile.Notes, Nickname: profile.Nickname}, nil
}

func (service *SimpleProfileService) ChangeBalance(value float32, profile *dto.Profile) error {
	newBalance := profile.Balance - float32(value)
	err := service.profileRepository.UpdateBalance(profile.Id, float64(newBalance))
	if err != nil {
		return fmt.Errorf("error changing balance: %w", err)
	}
	return nil
}
