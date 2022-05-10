package service

import (
	"playground.io/another-pet-store/db"
	"playground.io/another-pet-store/dto"
	"playground.io/another-pet-store/model"
)

type ChatService interface {
	CreateRoom(room *dto.ChatRoom) error
	GetRooms() (*dto.ChatRooms, error)
}

type SimpleChatService struct {
	chatRepository db.ChatRepository
}

func NewChatService(chatRepository db.ChatRepository) ChatService {
	return &SimpleChatService{chatRepository: chatRepository}
}

func (s *SimpleChatService) CreateRoom(room *dto.ChatRoom) error {
	r := s.chatRepository
	return r.CreateChatRoom(&model.ChatRoom{Name: room.Name})
}

func (s *SimpleChatService) GetRooms() (*dto.ChatRooms, error) {
	r := s.chatRepository
	rooms, err := r.FindAllRooms()
	if err != nil {
		return nil, err
	}
	var roomsDtoValues []dto.ChatRoom
	for _, v := range rooms.Rooms {
		roomsDtoValues = append(roomsDtoValues, dto.ChatRoom{ID: int(v.ID), Name: v.Name})
	}
	return &dto.ChatRooms{Rooms: roomsDtoValues}, nil
}
