package dto

type ChatRoom struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type ChatRooms struct {
	Rooms []ChatRoom `json:"rooms"`
}
