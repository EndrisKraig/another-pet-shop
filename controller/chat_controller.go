package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}

type ChatController interface {
	Chat(c *gin.Context)
}

type SimpleChatController struct {
}

func NewChatController() ChatController {
	return &SimpleChatController{}
}

func (controller *SimpleChatController) Chat(c *gin.Context) {
	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	//defer ws.Close()
	for {
		mt, message, err := ws.ReadMessage()
		if err != nil {
			fmt.Println(err)
			break
		}

		if string(message) == "ping" {
			message = []byte("pong")
		}

		err = ws.WriteMessage(mt, message)
		if err != nil {
			fmt.Println(err)
			break
		}
	}
}
