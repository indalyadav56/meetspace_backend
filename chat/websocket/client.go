package websocket

import (
	"io"
	"meetspace_backend/user/models"

	"github.com/gorilla/websocket"
)

type Client struct {
	Conn      *websocket.Conn
	Pool      *Pool
	GroupName string
	User      *models.User
	IsGroup   bool
}

func (c *Client) Read() {
	defer func() {
		c.Pool.Unregister <- c
		c.Conn.Close()
	}()

	for {
		_, msg, err := c.Conn.ReadMessage()

		if err != nil || err == io.EOF {
			c.Conn.WriteJSON(map[string]string{
				"error": "invalid payload",
			})
			break
		} else {
			c.Pool.Broadcast <- string(msg)
		}
	}
}
