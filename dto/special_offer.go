package dto

type SpecialOffers struct {
	Offers []SpecialOffer `json:"special_offers"`
}

type SpecialOffer struct {
	ID         int    `json:"id"`
	Nickname   string `json:"nickname"`
	Breed      string `json:"breed"`
	ImageUrl   string `json:"imageUrl"`
	Price      int    `json:"price"`
	BeginDate  string `json:"beginDate"`
	EndDate    string `json:"endDate"`
	Conditions string `json:"conditions"`
}
