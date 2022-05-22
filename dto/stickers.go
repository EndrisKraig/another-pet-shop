package dto

type Kits struct {
	Kits map[string][]Sticker `json:"kits"`
}

type Sticker struct {
	ID  int    `json:"id"`
	Uri string `json:"uri"`
}
