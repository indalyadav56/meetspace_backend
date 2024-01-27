package websocket

import (
	"github.com/gofiber/fiber/v2"
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

func handleWebSocketConnection(groupName string) fiber.Handler {
	return func(c *fiber.Ctx) error{
		currentPool := getOrCreatePool(groupName)
		WebSocketServer(currentPool, c)
		return nil
	}
}

func handleWebSocketConnectionFromParam() fiber.Handler {
	return func(c *fiber.Ctx) error{
		groupName := c.Query("groupName")
		currentPool := getOrCreatePool(groupName)
		WebSocketServer(currentPool, c)
		return nil
	}
}
