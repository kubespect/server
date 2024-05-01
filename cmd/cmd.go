package cmd

import (
	"log"

	"github.com/gin-gonic/gin"

	ws "github.com/kubespect/server/internal/websocket"
)

func Start() {

	hub := ws.NewHub()
	go hub.Run()

	//Start gin server
	app := gin.Default()

	app.GET("/ws/:roomId", func(c *gin.Context) {
		roomId := c.Param("roomId")
		ws.ServeWS(c, roomId, hub)

	})

	log.Println(app.Run(":8080"))

	// Get environment variables

	//Start websocket hub

	//Start grpc server

	//Listen for agent vitals

	//Retrieve data from alive agents

	//Send data to frontend via websocket

}
