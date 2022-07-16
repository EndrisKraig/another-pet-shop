package db

import (
	"context"
	"fmt"

	"gopkg.in/guregu/null.v4"
	"playground.io/another-pet-store/model"
)

type ProfileRepository interface {
	CreateProfile(userId int) error
	GetProfileByUserId(id int64) (*model.Profile, error)
	UpdateBalance(profileId int64, newBalance float64) error
}

type SimpleProfileRepository struct {
	connection Connection
}

func NewProfileRepository(connection Connection) ProfileRepository {
	return &SimpleProfileRepository{connection: connection}
}

func (r *SimpleProfileRepository) CreateProfile(userId int) error {
	conn, err := r.connection.GetConnection()

	if err != nil {
		return err
	}
	_, err = conn.Exec(context.Background(), "INSERT INTO user_profile (user_id, balance, image_url, notes) VALUES ($1, $2, $3, $4)", userId, 10000, "http://localhost:8080/", "Best user in the world!")

	if err != nil {
		return fmt.Errorf("error execute insert command: %w", err)
	}

	return nil
}

func (r *SimpleProfileRepository) GetProfileByUserId(ID int64) (*model.Profile, error) {
	conn, err := r.connection.GetConnection()

	if err != nil {
		return nil, err
	}

	var id int64
	var nickname string

	var image_url null.String
	var notes null.String
	var balance null.Float

	err = conn.QueryRow(context.Background(), "select user_profile.id, username, image_url, notes, balance from user_profile join users on user_id=users.id where user_id=$1", ID).Scan(&id, &nickname, &image_url, &notes, &balance)
	if err != nil {
		return nil, fmt.Errorf("error find profile with id %d: %w", id, err)
	}
	return &model.Profile{ID: id, Nickname: nickname, Image_url: image_url.ValueOrZero(), Notes: notes.ValueOrZero(), Balance: float32(balance.ValueOrZero())}, nil
}

func (r *SimpleProfileRepository) UpdateBalance(profileId int64, newBalance float64) error {
	conn, err := r.connection.GetConnection()

	if err != nil {
		return err
	}
	fmt.Println("New balance $1 for id $2", newBalance, profileId)
	_, err = conn.Exec(context.Background(), "UPDATE user_profile SET balance=$2 WHERE id = $1", profileId, newBalance)

	if err != nil {
		return fmt.Errorf("error execute update command: %w", err)
	}
	return nil
}
