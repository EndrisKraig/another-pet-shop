package service

import (
	"crypto/rand"
	"fmt"
)

type TicketService interface {
	CreateTicket(profileId int) string
	ReadTicket(ticket string) (int, error)
}

type SimpleTicketService struct {
	tickets map[string]int
}

func NewTicketService() TicketService {
	return &SimpleTicketService{tickets: make(map[string]int)}
}

func (s *SimpleTicketService) CreateTicket(profileId int) string {
	ticket := shortTicket(10)
	fmt.Println(ticket)
	s.tickets[ticket] = profileId
	return ticket
}

func (s *SimpleTicketService) ReadTicket(ticket string) (int, error) {
	profileId, ok := s.tickets[ticket]
	if !ok {
		return 0, fmt.Errorf("ticket %v not valid", ticket)
	}
	return profileId, nil
}

var chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"

func shortTicket(length int) string {
	fmt.Println("Gen start")
	ll := len(chars)
	b := make([]byte, length)
	rand.Read(b)
	for i := 0; i < length; i++ {
		b[i] = chars[int(b[i])%ll]
	}
	return string(b)
}
