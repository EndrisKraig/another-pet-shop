package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"playground.io/another-pet-store/dto"
	"playground.io/another-pet-store/service"
)

var upgrader = websocket.Upgrader{}

type ChatController interface {
	Chat(c *gin.Context)
	CreateRoom(c *gin.Context)
	GetAllRooms(c *gin.Context)
}

type SimpleChatController struct {
	chatService service.ChatService
}

func NewChatController(chatService service.ChatService) ChatController {
	return &SimpleChatController{chatService: chatService}
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

func (controller *SimpleChatController) CreateRoom(c *gin.Context) {
	var chatRoom dto.ChatRoom

	if err := c.BindJSON(&chatRoom); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Wrong body message"})
		return
	}

	err := controller.chatService.CreateRoom(&chatRoom)

	if err != nil {
		fmt.Println(err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Something went wrong :("})
		return
	}

	c.IndentedJSON(http.StatusCreated, gin.H{})
}

func (controller *SimpleChatController) GetAllRooms(c *gin.Context) {
	s := controller.chatService
	rooms, err := s.GetRooms()
	if err != nil {
		fmt.Println(err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Something went wrong :("})
		return
	}

	c.IndentedJSON(http.StatusOK, rooms)
}
