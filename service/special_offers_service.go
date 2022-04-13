package service

import (
	"fmt"

	"playground.io/another-pet-store/db"
	"playground.io/another-pet-store/dto"
)

type SpecialOfferService interface {
	GetAllActiveSpecialOffers() (*dto.SpecialOffers, error)
}

type SimpleSpecialOfferService struct {
	specialOfferRepository db.SpecialOfferRepository
}

func NewSpecialOfferServie(specialOfferRepository db.SpecialOfferRepository) SpecialOfferService {
	return &SimpleSpecialOfferService{specialOfferRepository: specialOfferRepository}
}

func (service *SimpleSpecialOfferService) GetAllActiveSpecialOffers() (*dto.SpecialOffers, error) {
	offersModel, err := service.specialOfferRepository.FindAllSpecialsOffers()
	if err != nil {
		return nil, fmt.Errorf("couldn't get special offers: %w", err)
	}
	var offers []dto.SpecialOffer
	for _, v := range offersModel.Offers {
		beginDate := v.BeginDate.String()
		endDate := v.EndDate.String()
		offers = append(offers, dto.SpecialOffer{ID: v.ID, Nickname: v.Nickname, Breed: v.Breed, Price: v.Price, BeginDate: beginDate, EndDate: endDate, Conditions: v.Conditions})
	}
	return &dto.SpecialOffers{Offers: offers}, nil
}
