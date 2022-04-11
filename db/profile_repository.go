package db

import (
	"context"
	"fmt"

	"playground.io/another-pet-store/model"
)

type ProfileRepository interface {
	CreateProfile(userId int) error
	GetProfileByUserId(id int64) (*model.Profile, error)
	UpdateBalance(profileId int64, newBalance float64) error
}

type SimpleProfileRepository struct{}

func NewProfileRepository() ProfileRepository {
	return &SimpleProfileRepository{}
}

func (rep *SimpleProfileRepository) CreateProfile(userId int) error {
	var conn = getConnection()
	defer conn.Close(context.Background())
	_, err := conn.Exec(context.Background(), "INSERT INTO user_profile (user_id, balance, image_url, notes) VALUES ($1, $2, $3, $4)", userId, 10000, "http://localhost:8080/", "Best user in the world!")

	if err != nil {
		return fmt.Errorf("error execute insert command: %w", err)
	}

	return nil
}

func (rep *SimpleProfileRepository) GetProfileByUserId(ID int64) (*model.Profile, error) {
	var conn = getConnection()
	defer conn.Close(context.Background())

	var id int64
	var nickname string

	var image_url string
	var notes string
	var balance float64

	err := conn.QueryRow(context.Background(), "select user_profile.id, username, image_url, notes, balance from user_profile join users on user_id=users.id where user_id=$1", ID).Scan(&id, &nickname, &image_url, &notes, &balance)
	if err != nil {
		return nil, fmt.Errorf("error find profile with id %d: %w", id, err)
	}
	return &model.Profile{ID: id, Nickname: nickname, Image_url: image_url, Notes: notes, Balance: float32(balance)}, nil
}

func (rep *SimpleProfileRepository) UpdateBalance(profileId int64, newBalance float64) error {
	var conn = getConnection()
	defer conn.Close(context.Background())
	_, err := conn.Exec(context.Background(), "UPDATE cats SET balance=$2 WHERE id = $1", profileId, newBalance)

	if err != nil {
		return fmt.Errorf("error execute update command: %w", err)
	}
	return nil
}
