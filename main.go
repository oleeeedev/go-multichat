package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"multichat/models"
)

func main() {
	router := gin.Default()

	router.Static("/static", "./static")

	router.GET("/", func(c *gin.Context) {
		c.File("./static/index.html")
	})

	router.POST("/send", func(c *gin.Context) {
		var message models.Message
		if err := c.ShouldBindJSON(&message); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		fmt.Printf("Nachricht von %s: %s\n", message.Username, message.Content)

		c.JSON(http.StatusOK, gin.H{"status": "message received"})
	})

	port := "8080"
	fmt.Printf("Server läuft auf http://localhost:%s\n", port)
	if err := router.Run(":" + port); err != nil {
		fmt.Printf("Fehler beim Starten des Servers: %s\n", err)
	}
}
