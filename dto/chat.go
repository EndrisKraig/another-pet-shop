package dto

import "time"

type ChatUser struct {
	ID      int
	Addr    string
	EnterAt time.Time
}

type Message struct {
	Sender int       `json:"sender"`
	Text   string    `json:"text"`
	SendAt time.Time `json:"sendAt"`
}

type Ticket struct {
	Ticket string `json:"ticket"`
}
