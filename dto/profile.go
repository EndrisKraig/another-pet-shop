package dto

type Profile struct {
	Id        int64   `json:"id"`
	Nickname  string  `json:"nickname"`
	Balance   float32 `json:"balance"`
	Image_url string  `json:"image_url"`
	Notes     string  `json:"Notes"`
}
