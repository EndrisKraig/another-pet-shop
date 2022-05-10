package service

import (
	"playground.io/another-pet-store/db"
	"playground.io/another-pet-store/dto"
	"playground.io/another-pet-store/model"
)

type MessageService interface {
	SaveMessage(message dto.Message, roomId int) error
	GetHistory(roomId int) ([]dto.Message, error)
}

type SimpleMessageServie struct {
	messageRepository db.MessageRepository
}

func NewMessageService(repository db.MessageRepository) MessageService {
	return &SimpleMessageServie{messageRepository: repository}
}

func (s *SimpleMessageServie) SaveMessage(message dto.Message, roomId int) error {
	return s.messageRepository.SaveMessage(&model.Message{Text: message.Text, ProfileId: message.Sender, RoomId: roomId})
}
func (s *SimpleMessageServie) GetHistory(roomId int) ([]dto.Message, error) {
	messages, err := s.messageRepository.GetHistory(roomId)
	if err != nil {
		return nil, err
	}
	dtoMessages := make([]dto.Message, 0)
	for _, v := range messages {
		dtoMessages = append(dtoMessages, dto.Message{Text: v.Text, Sender: v.ProfileId})
	}
	return dtoMessages, nil
}
