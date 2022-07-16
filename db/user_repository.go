package db

import (
	"context"
	"fmt"
	"os"

	"playground.io/another-pet-store/model"
)

type UserRepository interface {
	AddUser(user *model.User) error
	FindUser(username string) (*model.User, error)
}

type PostgresUserRepository struct {
	connection Connection
}

func NewUserRepository(conn Connection) UserRepository {
	return &PostgresUserRepository{connection: conn}
}

func (r *PostgresUserRepository) AddUser(user *model.User) error {
	conn, err := r.connection.GetConnection()

	if err != nil {
		return err
	}

	var id int64

	err = conn.QueryRow(context.Background(), "INSERT INTO users (username, pass_hash, email) VALUES ($1, $2, $3) RETURNING id", user.Username, user.Hash, user.Email).Scan(&id)

	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
	}

	err = conn.QueryRow(context.Background(), "INSERT INTO user_profile (nickname, user_id) VALUES ($1, $2) RETURNING id", user.Username, id).Scan(&id)

	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
	}
	return nil
}

func (r *PostgresUserRepository) FindUser(username string) (*model.User, error) {
	conn, err := r.connection.GetConnection()

	if err != nil {
		return nil, err
	}
	var id int64
	var hash string
	var email string
	err = conn.QueryRow(context.Background(), "select id, pass_hash, email FROM users WHERE username=$1", username).Scan(&id, &hash, &email)

	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	return &model.User{ID: id, Username: username, Hash: hash, Email: email}, nil

}
