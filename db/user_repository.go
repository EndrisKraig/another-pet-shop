package db

import (
	"context"
	"fmt"
	"os"

	"playground.io/another-pet-store/model"
)

func AddUser(user *model.User) {
	var conn = getConnection()
	defer conn.Close(context.Background())

	_, err := conn.Exec(context.Background(), "INSERT INTO users (username, pass_hash, email) VALUES ($1, $2, $3)", user.Username, user.Hash, user.Email)

	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}
}

func FindUser(username string) *model.User {
	var conn = getConnection()
	defer conn.Close(context.Background())
	var id int64
	var hash string
	var email string
	err := conn.QueryRow(context.Background(), "select id, pass_hash, email FROM users WHERE username=$1", username).Scan(&id, &hash, &email)

	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	return &model.User{ID: id, Username: username, Hash: hash, Email: email}

}
