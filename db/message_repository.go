package db

import (
	"context"
	"fmt"
	"time"

	"playground.io/another-pet-store/model"
)

type MessageRepository interface {
	SaveMessage(message *model.Message) error
	GetHistory(roomId int) ([]model.Message, error)
}

type SimpleMessageRepository struct {
}

func NewMessageRepository() MessageRepository {
	return new(SimpleMessageRepository)
}

func (t *SimpleMessageRepository) SaveMessage(message *model.Message) error {
	conn, err := GetConnection()

	if err != nil {
		return err
	}

	query := "INSERT INTO messages(profile_id, room_id, creation_date, text_body, send_status) VALUES ($1, $2, $3, $4, $5)"

	_, err = conn.Exec(context.Background(), query, message.ProfileId, message.RoomId, time.Now(), message.Text, 0)
	if err != nil {
		return err
	}

	return nil
}

func (t *SimpleMessageRepository) GetHistory(roomId int) ([]model.Message, error) {
	conn, err := GetConnection()

	if err != nil {
		return nil, err
	}

	query := "SELECT  id, profile_id, text_body FROM messages WHERE room_id = $1"

	rows, err := conn.Query(context.Background(), query, roomId)

	if err != nil {
		return nil, err
	}

	messages := make([]model.Message, 0)
	for rows.Next() {
		values, err := rows.Values()
		if err != nil {
			return nil, fmt.Errorf("error during obtaining result rows values: %w", err)
		}
		id := values[0].(int64)
		profile_id := values[1].(int)
		text := values[2].(string)

		messages = append(messages, model.Message{Id: int(id), ProfileId: profile_id, Text: text})

	}
	return messages, nil
}
