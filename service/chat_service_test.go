package service_test

import (
	"playground.io/another-pet-store/dto"
	"playground.io/another-pet-store/model"
	"playground.io/another-pet-store/service"
	"testing"
)

type StubChatRepository struct {
	rooms model.ChatRooms
	id    int
}

func (s *StubChatRepository) CreateChatRoom(room *model.ChatRoom) error {
	s.id++
	room.ID = int64(s.id)
	s.rooms.Rooms = append(s.rooms.Rooms, *room)
	return nil
}

func (s *StubChatRepository) FindAllRooms() (*model.ChatRooms, error) {
	return &s.rooms, nil
}

func TestName(t *testing.T) {
	chatService := service.NewChatService(&StubChatRepository{rooms: model.ChatRooms{Rooms: []model.ChatRoom{}}})
	err := chatService.CreateRoom(&dto.ChatRoom{Name: "MySuperRoom"})
	assertNoError(err, t)
	rooms, err := chatService.GetRooms()
	assertNoError(err, t)
	if rooms.Rooms == nil {
		t.Fatal("No rooms, but expected some")
	}
	if len(rooms.Rooms) != 1 {
		t.Errorf("Expecting exactly %d rooms, but got %d", 1, len(rooms.Rooms))
	}

}
