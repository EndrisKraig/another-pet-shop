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
	fmt.Print(message.Text)
	query := "INSERT INTO messages(profile_id, room_id, creation_date, text_body, send_status, format_id) VALUES ($1, $2, $3, $4, $5, (SELECT id FROM message_format WHERE label = $6))"

	_, err = conn.Exec(context.Background(), query, message.ProfileId, message.RoomId, time.Now(), message.Text, 0, message.Format)
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

	query := "SELECT  messages.id, profile_id, text_body, label FROM messages JOIN message_format ON messages.format_id = message_format.id WHERE room_id = $1"

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
		profile_id := values[1].(int32)
		text := values[2].(string)
		format := values[3].(string)
		messages = append(messages, model.Message{Id: int(id), ProfileId: int(profile_id), Text: text, Format: format})

	}
	return messages, nil
}
