package chat

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	"playground.io/another-pet-store/service"

	"github.com/gorilla/websocket"
)

const (
	// Time allowed to write a message to the peer.
	writeWait = 5 * time.Second

	// Time allowed to read the next pong message from the peer.
	pongWait = 5 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10

	// Maximum message size allowed from peer.
	maxMessageSize = 512
)

var (
	newline = []byte{'\n'}
	space   = []byte{' '}
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type User struct {
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
	Value string `json:"ticket"`
}

// Client is a middleman between the websocket connection and the hub.
type Client struct {
	hub           *Hub
	ticketService service.TicketService

	// The websocket connection.
	conn *websocket.Conn

	// Buffered channel of outbound messages.
	send chan Message
	//is ticket valid
	verified bool
	User
}

// readPump pumps messages from the websocket connection to the hub.
//
// The application runs readPump in a per-connection goroutine. The application
// ensures that there is at most one reader on a connection by executing all
// reads from this goroutine.
func (c *Client) readPump() {
	defer func() {
		c.hub.unregister <- c
		c.conn.Close()
	}()
	c.conn.SetReadLimit(maxMessageSize)
	c.conn.SetReadDeadline(time.Now().Add(pongWait))
	c.conn.SetPongHandler(func(string) error { c.conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	for {
		_, p, err := c.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}
		if c.verified {
			var data *Message = &Message{}
			err = json.Unmarshal(p, data)
			data.Sender = c.ID
			if err != nil {
				fmt.Println(err)
				return
			}

			c.hub.broadcast <- *data
		} else {
			var ticket *Ticket = new(Ticket)
			err = json.Unmarshal(p, ticket)

			if err != nil {
				fmt.Println(err)
				return
			}

			ticketService := c.ticketService
			profileId, err := ticketService.ReadTicket(ticket.Value)
			if err != nil {
				c.hub.unregister <- c
				c.conn.Close()
				fmt.Println(err)
			}

			c.ID = profileId
			c.verified = true
			fmt.Println("Registered")
			c.hub.register <- c
		}
	}
}

// writePump pumps messages from the hub to the websocket connection.
//
// A goroutine running writePump is started for each connection. The
// application ensures that there is at most one writer to a connection by
// executing all writes from this goroutine.
func (c *Client) writePump() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.conn.Close()
	}()
	for {
		select {
		case message, ok := <-c.send:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				// The hub closed the channel.
				c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			w, err := c.conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}

			message.SendAt = time.Now()
			jsonText, _ := json.Marshal(message)

			w.Write(jsonText)

			// Add queued chat messages to the current websocket message.
			n := len(c.send)
			for i := 0; i < n; i++ {
				w.Write(newline)
				data, _ := json.Marshal(<-c.send)
				w.Write(data)
			}

			if err := w.Close(); err != nil {
				return
			}
		case <-ticker.C:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}

// serveWs handles websocket requests from the peer.
func ServeWs(hub *Hub, w http.ResponseWriter, r *http.Request, ticketService service.TicketService) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	client := &Client{hub: hub, ticketService: ticketService, conn: conn, send: make(chan Message, 256), verified: false}
	client.hub.register <- client
	client.Addr = conn.RemoteAddr().String()
	client.EnterAt = time.Now()

	// Allow collection of memory referenced by the caller by doing all work in
	// new goroutines.
	go client.writePump()
	go client.readPump()

	//client.send <- []byte("Welcome")
}

func GenUserId() string {
	uid := uuid.NewString()
	return uid
}
