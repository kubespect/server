package cmd

import (
	"log"

	"github.com/gin-gonic/gin"

	"github.com/kubespect/server/internal/grpc"
	ws "github.com/kubespect/server/internal/websocket"
)

func Start() {
	grpcServer := grpc.NewGrpc()
	go grpcServer.Start()

	hub := ws.NewHub()
	go hub.Run()
	// gin.SetMode(gin.ReleaseMode)
	// Start gin server.
	app := gin.Default()

	app.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World",
		})
	})
	app.GET("/ws/:roomId", func(c *gin.Context) {
		roomID := c.Param("roomID")
		ws.ServeWS(c, roomID, hub)
	})

	log.Println(app.Run(":8080"))

	// Get environment variables.

	// Start websocket hub.

	// Start grpc server.

	// Listen for agent vitals.

	// Retrieve data from alive agents.

	// Send data to frontend via websocket.
}
