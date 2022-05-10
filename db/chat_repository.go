package db

import (
	"context"

	"playground.io/another-pet-store/model"
)

type ChatRepository interface {
	CreateChatRoom(room *model.ChatRoom) error
	FindAllRooms() (*model.ChatRooms, error)
}

type SimpleChatRepository struct {
}

func NewChatRepository() ChatRepository {
	return new(SimpleChatRepository)
}

func (r *SimpleChatRepository) CreateChatRoom(room *model.ChatRoom) error {
	conn, err := GetConnection()

	if err != nil {
		return err
	}

	query := "INSERT INTO conversation_room(room_type, room_name) VALUES ($1, $2)"

	_, err = conn.Exec(context.Background(), query, 0, room.Name)

	return err
}

func (r *SimpleChatRepository) FindAllRooms() (*model.ChatRooms, error) {
	conn, err := GetConnection()

	if err != nil {
		return nil, err
	}

	query := "SELECT id, room_type, room_name FROM conversation_room"

	rows, err := conn.Query(context.Background(), query)

	if err != nil {
		return nil, err
	}

	var rooms []model.ChatRoom

	for rows.Next() {
		values, err := rows.Values()
		if err != nil {
			return nil, err
		}
		id := values[0].(int64)
		name := values[2].(string)
		rooms = append(rooms, model.ChatRoom{ID: id, Name: name})

	}
	return &model.ChatRooms{Rooms: rooms}, nil
}
