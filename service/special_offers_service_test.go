package service_test

import (
	"fmt"
	"testing"
	"time"

	"playground.io/another-pet-store/dto"
	"playground.io/another-pet-store/model"
	"playground.io/another-pet-store/service"
)

type StubSpecialOfferRepository struct {
	specialOffers *model.SpecialOffers
}

func (r *StubSpecialOfferRepository) FindAllSpecialsOffers() (*model.SpecialOffers, error) {
	return r.specialOffers, nil
}

var idGenerator int

func createSpecialOffer(nickname, breed string, begin, end time.Time) *model.SpecialOffer {
	idGenerator++
	return &model.SpecialOffer{ID: idGenerator, Nickname: nickname, Breed: breed, ImageUrl: fmt.Sprintf("http://specials.io/%d", idGenerator), BeginDate: begin, EndDate: end}
}

func TestSpecialOffers(t *testing.T) {
	offers := []model.SpecialOffer{*createSpecialOffer("Faeron", "persian", time.Now(), time.Now().Add(time.Hour))}
	specialOffers := model.SpecialOffers{Offers: offers}
	specialOffersService := service.NewSpecialOfferService(&StubSpecialOfferRepository{specialOffers: &specialOffers})
	resp, err := specialOffersService.GetAllActiveSpecialOffers()
	assertNoError(err, t)
	assertNotNil(resp, t)
	getOffers := resp.Offers
	if len(getOffers) != 1 {
		t.Errorf("Expected exactly one offer, but got %d", len(getOffers))
	}
}

func assertNotNil(resp *dto.SpecialOffers, t *testing.T) {
	t.Helper()
	if resp == nil {
		t.Errorf("Expected value, but got nil")
	}
}
