package websocket

import (
	"fmt"
	"log"
	"net/http"

	"meetspace_backend/utils"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func Upgrade(reqWriter http.ResponseWriter, req *http.Request) (*websocket.Conn, error) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	conn, err := upgrader.Upgrade(reqWriter, req, nil)
	
	if err != nil {
		log.Println("Websocket connection error:- ", err)
		return nil, err
	}
	
	return conn, nil
}

func WebSocketServer(pool *Pool, c *gin.Context)  {
	conn, err := Upgrade(c.Writer, c.Request)
	
	if err != nil {
		fmt.Println("WebSocket connection error: ", err)
		return
	}
	
	groupName := c.Param("groupName")
	
	currentUser, exists := utils.GetUserFromContext(c)

	if !exists{
		return 
	}

	client := &Client{
		Conn:      conn,
		Pool:      pool,
		GroupName: groupName,
		IsGroup:   groupName != "",
		User:      currentUser,
	}

	pool.Register <- client
	client.Read()
}
