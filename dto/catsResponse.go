package dto

type CatsResponse struct {
	Cats    []Cat `json:"cats"`
	MaxPage int   `json:"maxPage"`
}

type Cat struct {
	ID       int64  `json:"id"`
	Nickname string `json:"nickname"`
	Breed    string `json:"breed"`
	Price    int32  `json:"price"`
}
