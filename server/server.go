package server

import (
	"github.com/gin-gonic/gin"
	"multichat/models"
)

func main() {
	r := gin.Default()

	r.POST("/register", func(c *gin.Context) {
		var user models.User
		if err := c.ShouldBindJSON(&user); err == nil {
			client := &server.Client{Username: user.Username, Conn: c}
			server.Clients[user.Username] = client
			c.JSON(200, gin.H{"status": "registered"})
		} else {
			c.JSON(400, gin.H{"error": "bad request"})
		}
	})

	r.POST("/chat/:room", func(c *gin.Context) {
		roomName := c.Param("room")
		var message models.Message
		if err := c.ShouldBindJSON(&message); err == nil {
			server.SendMessage(roomName, message)
			c.JSON(200, gin.H{"status": "message sent"})
		} else {
			c.JSON(400, gin.H{"error": "bad request"})
		}
	})

	r.GET("/chat/:room", func(c *gin.Context) {
		roomName := c.Param("room")
		server.CreateChatRoom(roomName)
		client := &server.Client{Username: c.Query("username"), Conn: c}
		server.JoinChatRoom(roomName, client)
	})

	r.Run(":8080")
}
