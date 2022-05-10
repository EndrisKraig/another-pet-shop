package chat

import (
	"fmt"
	"strconv"
	"time"

	"playground.io/another-pet-store/dto"
	"playground.io/another-pet-store/service"
)

//credit for https://hoohoo.top/blog/20220320172715-go-websocket/

// Hub maintains the set of active clients and broadcasts messages to the
// clients.
type Hub struct {
	//aka room id
	id int
	// Registered clients.
	clients map[*Client]bool

	// Inbound messages from the clients.
	broadcast chan dto.Message

	// Register requests from the clients.
	register chan *Client

	// Unregister requests from clients.
	unregister chan *Client

	messageService service.MessageService
}

func NewHub(id int, messageService service.MessageService) *Hub {
	return &Hub{
		id:             id,
		broadcast:      make(chan dto.Message),
		register:       make(chan *Client),
		unregister:     make(chan *Client),
		clients:        make(map[*Client]bool),
		messageService: messageService,
	}
}

type History struct {
	Type     string        `json:"type"`
	Messages []dto.Message `json:"messages"`
}

func (h *Hub) Run() {
	messageService := h.messageService
	for {
		select {
		case client := <-h.register:
			_, ok := h.clients[client]
			if client.verified && !ok {
				messages, err := messageService.GetHistory(h.id)
				if err != nil {
					fmt.Println(err)
				}
				history := History{Messages: messages, Type: "history"}
				client.sendHistory(history)
				clientId := client.ID
				for client := range h.clients {
					msg := dto.Message{Sender: 777, SendAt: time.Now(), Text: "Some one connected: " + strconv.Itoa(clientId)}
					client.send <- msg
				}

				h.clients[client] = true
			}

		case client := <-h.unregister:
			clientId := client.ID
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.send)
			}
			for client := range h.clients {
				msg := dto.Message{Sender: 777, SendAt: time.Now(), Text: "Some one leave: " + strconv.Itoa(clientId)}
				client.send <- msg
			}
		case userMessage := <-h.broadcast:
			messageService.SaveMessage(userMessage, h.id)
			for client := range h.clients {
				//prevent self receive the message
				if client.ID == userMessage.Sender {
					continue
				}
				select {
				case client.send <- userMessage:
				default:
					close(client.send)
					delete(h.clients, client)
				}
			}
		}
	}
}
