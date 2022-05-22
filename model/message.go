package model

import (
	"time"
)

type Message struct {
	Id           int
	ProfileId    int
	RoomId       int
	CreationDate time.Time
	Text         string
	SendStatus   int
	Format       string
}
