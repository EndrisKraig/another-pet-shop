package dto

type Reference struct {
	Id    int    `json:"id"`
	Label string `json:"label"`
}

type References struct {
	References []Reference `json:"references"`
}
