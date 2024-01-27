package websocket

import (
	"github.com/gin-gonic/gin"
)

// gets or create WebSocket pool based on the groupName
func getOrCreatePool(groupName string) *Pool {
	pool, exists := groupPools[groupName]
	if !exists {
		pool = NewPool()
		groupPools[groupName] = pool
		go pool.Start()
	}
	return pool
}

func handleWebSocketConnection(groupName string) gin.HandlerFunc {
	return func(c *gin.Context) {
		currentPool := getOrCreatePool(groupName)
		WebSocketServer(currentPool, c)
	}
}

func handleWebSocketConnectionFromParam() gin.HandlerFunc {
	return func(c *gin.Context) {
		groupName := c.Param("groupName")
		currentPool := getOrCreatePool(groupName)
		WebSocketServer(currentPool, c)
	}
}