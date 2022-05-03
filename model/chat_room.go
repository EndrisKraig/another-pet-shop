package model

type ChatRoom struct {
	ID   int64
	Name string
}

type ChatRooms struct {
	Rooms []ChatRoom
}
