package db

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v4"
)

type Connection interface {
	GetConnection() (*pgx.Conn, error)
}

type PgConnection struct {
	connection *pgx.Conn
}

func (c *PgConnection) GetConnection() (*pgx.Conn, error) {
	if c.connection == nil {
		conn, err := pgx.Connect(context.Background(), os.Getenv("DB_URL"))
		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
			os.Exit(1)
		}
		c.connection = conn
	}
	err := c.connection.Ping(context.Background())
	if err != nil {
		conn, err := pgx.Connect(context.Background(), os.Getenv("DB_URL"))
		if err != nil {
			return nil, fmt.Errorf("Unable to connect to database: %v\n", err)
		}
		c.connection = conn
	}
	return c.connection, nil
}
