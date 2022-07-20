package service_test

import (
	"testing"

	"playground.io/another-pet-store/dto"
	"playground.io/another-pet-store/model"
	"playground.io/another-pet-store/service"
)

type StubMessageRepository struct {
	messages []model.Message
}

var roomId = 1

func (r *StubMessageRepository) SaveMessage(message *model.Message) error {
	r.messages = append(r.messages, *message)
	return nil
}
func (r *StubMessageRepository) GetHistory(roomId int) ([]model.Message, error) {
	return r.messages, nil
}

func TestMessage(t *testing.T) {
	repository := StubMessageRepository{
		[]model.Message{
			model.Message{Id: 1, Text: "Hello!", ProfileId: 1},
			model.Message{Id: 1, Text: "World?", ProfileId: 2},
			model.Message{Id: 1, Text: "No, just greeting...", ProfileId: 1},
		},
	}
	messageService := service.NewMessageService(&repository)

	history, err := messageService.GetHistory(1)
	assertNoError(err, t)

	n := len(history)

	if n != 3 {
		t.Errorf("Expected %d messages, got %d", 3, n)
	}

	err = messageService.SaveMessage(dto.Message{}, 1)
	assertNoError(err, t)

	history, err = messageService.GetHistory(1)
	assertNoError(err, t)

	n = len(history)

	if n != 4 {
		t.Errorf("Expected %d messages, got %d", 4, n)
	}

}
