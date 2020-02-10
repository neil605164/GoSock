package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

func main() {

	r := gin.Default()

	// r.GET("/", func(c *gin.Context) {
	// 	c.String(200, "We got Gin")
	// })

	r.LoadHTMLFiles("index.html")

	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", nil)
	})

	r.GET("/ws", wshandler)

	r.Run()
}

var wsupgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func wshandler(c *gin.Context) {
	conn, err := wsupgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println("Failed to set websocket upgrade: ", err)
		return
	}

	for {
		t, msg, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("Error Msg ===>", err.Error())
			break
		}

		fmt.Println(t)
		fmt.Println(string(msg))
		conn.WriteMessage(t, msg)
	}
}
