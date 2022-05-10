package model

import "time"

type SpecialOffers struct {
	Offers []SpecialOffer
}

type SpecialOffer struct {
	ID         int
	Nickname   string
	Breed      string
	ImageUrl   string
	Price      int
	BeginDate  time.Time
	EndDate    time.Time
	Conditions string
}
