package server

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sync"

	"multichat/models"
)

type Client struct {
	Username string
	Conn     *gin.Context
}

type ChatRoom struct {
	Name    string
	Clients map[*Client]bool
	Mutex   sync.Mutex
}

var (
	ChatRooms    = make(map[string]*ChatRoom)
	Clients      = make(map[string]*Client)
	ClientsMutex = sync.Mutex{}
)

func CreateChatRoom(name string) {
	ChatRooms[name] = &ChatRoom{Name: name, Clients: make(map[*Client]bool)}
}

func JoinChatRoom(roomName string, client *Client) {
	ClientsMutex.Lock()
	Clients[client.Username] = client
	ClientsMutex.Unlock()

	ChatRooms[roomName].Mutex.Lock()
	ChatRooms[roomName].Clients[client] = true
	ChatRooms[roomName].Mutex.Unlock()
}

func SendMessage(roomName string, message models.Message) {
	ChatRooms[roomName].Mutex.Lock()
	defer ChatRooms[roomName].Mutex.Unlock()

	for client := range ChatRooms[roomName].Clients {
		client.Conn.JSON(http.StatusOK, message)
	}
}
