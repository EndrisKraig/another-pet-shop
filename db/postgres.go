package db

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v4"
)

var connection *pgx.Conn

func GetConnection() (*pgx.Conn, error) {
	if connection == nil {
		conn, err := pgx.Connect(context.Background(), os.Getenv("DB_URL"))
		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
			os.Exit(1)
		}
		connection = conn
	}
	err := connection.Ping(context.Background())
	if err != nil {
		conn, err := pgx.Connect(context.Background(), os.Getenv("DB_URL"))
		if err != nil {
			return nil, fmt.Errorf("Unable to connect to database: %v\n", err)
		}
		connection = conn
	}
	return connection, nil
}
