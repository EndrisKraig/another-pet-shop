package model

type Cat struct {
	ID       int32  `json:"id"`
	Nickname string `json:"nickname"`
	Breed    string `json:"breed"`
	Price    int32  `json:"price"`
}
