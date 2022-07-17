package service_test

import (
	"fmt"
	"testing"

	"playground.io/another-pet-store/model"
	"playground.io/another-pet-store/service"
)

type StubProfileRepository struct {
	profiles map[int]model.Profile
}

func (r StubProfileRepository) CreateProfile(userId int) error {
	r.profiles[userId] = model.Profile{ID: int64(userId)}
	return nil
}
func (r StubProfileRepository) GetProfileByUserId(id int64) (*model.Profile, error) {
	profile, ok := r.profiles[int(id)]
	if !ok {
		return nil, fmt.Errorf("No profile with id: %d", id)
	}
	return &profile, nil
}
func (r StubProfileRepository) UpdateBalance(profileId int64, newBalance float64) error {
	return nil
}

func TestBalanceChange(t *testing.T) {
	profileService := createProfileService()

	balance, err := profileService.ChangeBalance(5, 10)
	want := 5

	assertNoError(err, t)

	if balance != 5 {
		t.Errorf("Balance is %q, wanted %q", balance, want)
	}

	_, err = profileService.ChangeBalance(15, 10)

	assertError(err, t)
}

func createProfileService() service.ProfileService {
	testProfiles := make(map[int]model.Profile)
	testProfiles[1] = model.Profile{ID: 1, Balance: 1000}
	profileService := service.NewProfileService(StubProfileRepository{testProfiles})
	return profileService
}

func TestCreateGetProfile(t *testing.T) {
	profileService := createProfileService()
	err := profileService.CreateProfile(2)
	assertNoError(err, t)
	profile, err := profileService.GetProfile(2)
	assertNoError(err, t)
	if profile == nil {
		t.Fatalf("Wanted profile, but got nil")
	}

	if profile.Id != 2 {
		t.Errorf("Wanted %d, but got %d", 2, profile.Id)
	}

	_, err = profileService.GetProfile(42)
	assertError(err, t)

}

func TestNoEnoughBalance(t *testing.T) {

}
